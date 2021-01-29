package chess

type Game struct {
	*Board
	PlayedMoves []*Move

	legalMoves       []*Move
	startingPosition *Position
}

func NewGame() *Game {
	return NewGameWithPosition(MustParseFenPosition(StartingPosition))
}

func NewGameWithPosition(position *Position) *Game {
	return (&Game{
		Board:            position.Placement.Board(),
		startingPosition: position,
	}).updateLegalMoves()
}

func NewGameWithPlacement(placement Placement) *Game {
	return NewGameWithPosition(NewPosition(placement, White))
}

func LoadGame(moves []MoveRequest) (*Game, error) {
	game := NewGame()
	for _, req := range moves {
		if _, err := game.Move(req); err != nil {
			return nil, err
		}
	}
	return game, nil
}

func (g *Game) IsStarted() bool {
	return len(g.PlayedMoves) > 0
}

func (g *Game) NextPlayer() Color {
	if !g.IsStarted() {
		return g.startingPosition.NextPlayer
	}
	return !g.lastMove().Color
}

func (g *Game) LegalMoves() []*Move {
	return g.legalMoves
}

func (g *Game) Validate(req MoveRequest) (*Move, error) {
	if g.IsEnded() {
		return nil, ErrGameHasEnded
	}
	for _, move := range g.legalMoves {
		if move.MoveRequest == req {
			return move, nil
		}
	}
	return nil, ErrMoveNotAllowed
}

func (g *Game) Move(req MoveRequest) (move *Move, err error) {
	move, err = g.Validate(req)
	if err != nil {
		return
	}
	g.applyMove(move)
	g.PlayedMoves = append(g.PlayedMoves, move)
	g.updateLegalMoves()
	return
}

func (g *Game) MustMove(req MoveRequest) *Move {
	move, err := g.Move(req)
	if err != nil {
		panic(err)
	}
	return move
}

func (g *Game) TakeBack() error {
	if !g.IsStarted() {
		return ErrNoFirstMove
	}
	g.revertMove(g.lastMove())
	g.PlayedMoves = g.PlayedMoves[:len(g.PlayedMoves)-1]
	g.updateLegalMoves()
	return nil
}

func (g *Game) HalfMoveClock() int {
	last := len(g.PlayedMoves) - 1
	for i := range g.PlayedMoves {
		move := g.PlayedMoves[last-i]
		if move.IsCapture() || move.Type == Pawn {
			return i
		}
	}
	return g.startingPosition.HalfMoves + len(g.PlayedMoves)
}

func (g *Game) FullMoves() int {
	return g.startingPosition.FullMoves + len(g.PlayedMoves)/2
}

func (g *Game) Position() *Position {
	return &Position{
		Placement:            g.Placement(),
		NextPlayer:           g.NextPlayer(),
		CastlingAvailability: g.castlingAvailability(),
		EnPassantTarget:      g.enPassantTarget(),
		HalfMoves:            g.HalfMoveClock(),
		FullMoves:            g.FullMoves(),
	}
}

func (g *Game) applyMove(move *Move) {
	if move.IsCapture() {
		g.Board.At(move.CapturedPieceCoordinate()).Piece = nil
	}
	g.Board.At(move.To).Piece = move.Piece
	g.Board.At(move.From).Piece = nil
	if move.IsPromotion() {
		move.Piece.Type = move.PromoteTo
	}
	if move.IsCastling() {
		req := move.CastlingMoveRequest()
		g.applyMove(&Move{
			MoveRequest: req,
			Piece:       g.At(req.From).Piece,
		})
	}
}

func (g *Game) revertMove(move *Move) {
	if move.IsCastling() {
		req := move.CastlingMoveRequest()
		g.revertMove(&Move{
			MoveRequest: req,
			Piece:       g.At(req.To).Piece,
		})
	}
	if move.IsPromotion() {
		move.Piece.Type = Pawn
	}
	g.Board.At(move.From).Piece = move.Piece
	g.Board.At(move.To).Piece = nil
	if move.IsCapture() {
		g.Board.At(move.CapturedPieceCoordinate()).Piece = move.CapturedPiece
	}
}

func (g *Game) lastMove() *Move {
	if !g.IsStarted() {
		return nil
	}
	return g.PlayedMoves[len(g.PlayedMoves)-1]
}
