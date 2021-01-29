package chess

import "sort"

var (
	// Scan patterns start north (+1 rank) and go clockwise
	RookVectors   = C3.VectorsTo(C4, D3, C2, B3)
	BishopVectors = C3.VectorsTo(D4, D2, B2, B4)
	KnightVectors = C3.VectorsTo(D5, E4, E2, D1, B1, A2, A4, B5)
	RoyalVectors  = C3.VectorsTo(C4, D4, D3, D2, C2, B2, B3, B4)
)

func (g *Game) updateLegalMoves() *Game {
	g.legalMoves = g.findLegalMoves()
	return g
}

func (g *Game) findLegalMoves() (moves []*Move) {
	var (
		player       = g.NextPlayer()
		opposingKing = g.findKing(!player)
	)
	for _, from := range g.Board {
		if from.IsVacant() || from.Piece.Color != player {
			continue
		}
		for _, to := range g.reachableSquares(from, false) {
			move := &Move{
				MoveRequest: MoveRequest{
					From: from.Coordinate,
					To:   to.Coordinate,
				},
				Piece:         from.Piece,
				CapturedPiece: to.Piece,
			}

			// En passant
			if from.Piece.Type == Pawn && to.IsVacant() && from.File() != to.File() {
				move.IsEnPassant = true
				move.CapturedPiece = g.At(NewCoordinate(to.File(), from.Rank())).Piece
			}

			candidates := []*Move{move}

			// Promotion
			if from.Piece.Type == Pawn && to.Rank()%7 == 0 {
				candidates = make([]*Move, 4)
				for i, pieceType := range []PieceType{Queen, Knight, Rook, Bishop} {
					m := *move
					m.MoveRequest.PromoteTo = pieceType
					candidates[i] = &m
				}
			}

			for _, move := range candidates {
				// Look-ahead
				g.applyMove(move)

				// Disallow moves resulting in self-check
				if !g.isAttacked(g.findKing(player), !player) {

					// Mark moves resulting in check
					if g.isAttacked(opposingKing, player) {
						move.IsCheck = true
					}

					moves = append(moves, move)
				}

				// Undo look-ahead
				g.revertMove(move)
			}
		}
	}
	return
}

// Does not cover en passant attacks; used for testing check and clearing castling moves
func (g *Game) isAttacked(square *Square, byColor Color) bool {
	for _, from := range g.Board {
		if from.IsVacant() || from.Piece.Color != byColor {
			continue
		}
		for _, to := range g.reachableSquares(from, true) {
			if to == square {
				return true
			}
		}
	}
	return false
}

func (g *Game) isInsufficientMaterial() bool {
	remainingPieces := make(map[Color][]*Square)
	for _, square := range g.Board {
		if square.IsOccupied() {
			remainingPieces[square.Piece.Color] = append(remainingPieces[square.Piece.Color], square)
		}
	}
	for _, squares := range remainingPieces {
		sort.Slice(squares, func(i, j int) bool {
			return squares[i].Piece.Type < squares[j].Piece.Type
		})
	}
	for sideA, remainingA := range remainingPieces {
		var (
			sideB      = !sideA
			remainingB = remainingPieces[sideB]
		)
		switch len(remainingA) {
		// TODO: bishops with same FG and BG color (underpromotion)
		case 1:
			switch len(remainingB) {
			case 1: // King vs King
				return true
			case 2: // King vs King and Bishop or Knight
				t := remainingB[0].Piece.Type
				return t == Bishop || t == Knight
			}
		case 2:
			// King and Bishop vs King and Bishop on same color
			if remainingA[0].Piece.Type == Bishop &&
				len(remainingB) == 2 &&
				remainingB[0].Piece.Type == Bishop &&
				remainingA[0].BackgroundColor() == remainingB[0].BackgroundColor() {
				return true
			}
		}
	}
	return false
}

func pawnAdvanceDirection(color Color) int {
	if color == White {
		return 1
	}
	return -1
}

func pawnHomeRank(color Color) int {
	if color == White {
		return 1
	}
	return 6
}

func (g *Game) reachableSquares(from *Square, excludeCastling bool) (squares []*Square) {
	switch from.Piece.Type {
	case Pawn:
		var (
			home     = pawnHomeRank(from.Piece.Color)
			dir      = pawnAdvanceDirection(from.Piece.Color)
			limit    = 1
			fromRank = from.Rank()
		)
		if fromRank == home {
			limit = 2
		}
		squares = append(
			g.scan(scan{from: from, vectors: []Vector{{0, dir}}, limit: limit, excludeCaptures: true}),
			g.scan(scan{from: from, vectors: []Vector{{1, dir}, {-1, dir}}, limit: 1, onlyCaptures: true})...,
		)
	case Knight:
		squares = g.scan(scan{from: from, vectors: KnightVectors, limit: 1})
	case Bishop:
		squares = g.scan(scan{from: from, vectors: BishopVectors})
	case Rook:
		squares = g.scan(scan{from: from, vectors: RookVectors})
	case Queen:
		squares = g.scan(scan{from: from, vectors: RoyalVectors})
	case King:
		squares = g.scan(scan{from: from, vectors: RoyalVectors, limit: 1})
		if !excludeCastling {
			squares = append(squares, g.reachableSquaresByCastling(from)...)
		}
	}
	return
}

func (g *Game) enPassantTarget() Coordinate {
	lastMove := g.lastMove()
	if lastMove == nil {
		return g.startingPosition.EnPassantTarget
	}
	var (
		home = pawnHomeRank(lastMove.Color)
		dir  = pawnAdvanceDirection(lastMove.Color)
	)
	if lastMove.Type == Pawn &&
		lastMove.From.Rank() == home &&
		lastMove.To.Rank() == home+dir*2 {
		return NewCoordinate(lastMove.From.File(), home+dir)
	}
	return 0
}

func (g *Game) castlingAvailability() CastlingAvailability {
	var (
		a          = g.startingPosition.CastlingAvailability.Copy()
		foundKings = make(map[Color]bool)
	)
	for _, move := range g.PlayedMoves {
		if len(a) == 0 {
			return a
		}
		switch move.Type {
		case Rook:
			a.Unset(move.From)
		case King:
			if !foundKings[move.Color] {
				rank := move.From.Rank()
				for file := 0; file <= 7; file++ {
					a.Unset(NewCoordinate(file, rank))
				}
				foundKings[move.Color] = true
			}
		}
	}
	return a
}

type castlingSide struct {
	path     []int
	rookFile int
}

var castlingSides = [2]castlingSide{
	{[]int{5, 6}, 7},    // Kingside
	{[]int{3, 2, 1}, 0}, // Queenside
}

func (g *Game) reachableSquaresByCastling(kingSquare *Square) (squares []*Square) {
	if g.IsCheck() || kingSquare.File() != 4 || g.pieceHasMoved(kingSquare.Piece) {
		return
	}

	var (
		rank         = kingSquare.Rank()
		availability = g.castlingAvailability() // TODO: vary to only search for one color
	)

	for _, opt := range castlingSides {
		// Make sure the expected rook is in place and hasn't moved
		if !availability[NewCoordinate(opt.rookFile, rank)] {
			continue
		}

		// Make sure there are no pieces between the king and rook
		obstructed := false
		for _, file := range opt.path {
			if g.At(NewCoordinate(file, rank)).IsOccupied() {
				obstructed = true
				break
			}
		}
		if obstructed {
			continue
		}

		// Make sure the king is not passing through an attacked square
		if g.isAttacked(g.At(NewCoordinate(opt.path[0], rank)), !kingSquare.Piece.Color) {
			continue
		}

		squares = append(squares, g.At(NewCoordinate(opt.path[1], rank)))
	}
	return
}

type scan struct {
	from            *Square
	limit           int
	excludeCaptures bool
	onlyCaptures    bool
	vectors         []Vector
}

func (g *Game) scan(s scan) (matches []*Square) {
	if s.limit == 0 {
		s.limit = 7
	}
	for _, vector := range s.vectors {
		for distance := 1; distance <= s.limit; distance += 1 {
			toCoord, err := s.from.Coordinate.Shift(vector, distance)
			if err != nil {
				continue
			}
			to := g.Board.At(toCoord)
			if to.IsOccupied() {
				if !s.excludeCaptures && to.Piece.Color != s.from.Piece.Color {
					matches = append(matches, to)
				}
				break
			} else if s.from.Piece.Type == Pawn && to.Coordinate != 0 && to.Coordinate == g.enPassantTarget() {
				matches = append(matches, to)
			}
			if !s.onlyCaptures {
				matches = append(matches, to)
			}
		}
	}
	return
}

// Does not report rooks moved by castling. Used for determining validity of castling and en passant.
func (g *Game) pieceHasMoved(piece *Piece) bool {
	for _, move := range g.PlayedMoves {
		if move.Piece == piece {
			return true
		}
	}
	return false
}

func (g *Game) findKing(color Color) (square *Square) {
	for _, square = range g.Board {
		if square.Piece.Is(color, King) {
			return
		}
	}
	panic("game board does not have a " + color.String() + " King")
}
