package chess

import (
	"fmt"
	"strconv"
	"strings"
)

func (b *Board) Fen() string {
	var (
		seq    = make([]byte, 0, 15)
		blanks byte
	)
	var flush = func() {
		if blanks > 0 {
			seq = append(seq, '0'+blanks)
			blanks = 0
		}
	}
	for rank := 7; rank >= 0; rank-- {
		if rank != 7 {
			seq = append(seq, '/')
		}
		for file := 0; file <= 7; file++ {
			piece := b.At(NewCoordinate(file, rank)).Piece
			if piece == nil {
				blanks++
			} else {
				flush()
				seq = append(seq, byte(piece.Letter()))
			}
		}
		flush()
	}
	return string(seq)
}

func (p Placement) Fen() string {
	return p.Board().Fen()
}

func (g *Game) Fen() string {
	return g.Position().Fen()
}

func (a CastlingAvailability) Fen() string {
	shredder := false
	for c := range a {
		if c.File()%7 != 0 {
			shredder = true
			break
		}
	}
	var result []byte
	for i := range [2]struct{}{} { // White = 0, Black = 1
		var (
			rank  = i * 7
			lower = byte(i) * 32
		)
		if shredder {
			for file := 0; file <= 7; file++ {
				if a[NewCoordinate(file, rank)] {
					result = append(result, 'A'+byte(file)+lower)
				}
			}
		} else {
			if a[NewCoordinate(7, rank)] {
				result = append(result, 'K'+lower)
			}
			if a[NewCoordinate(0, rank)] {
				result = append(result, 'Q'+lower)
			}
		}
	}

	if len(result) == 0 {
		return "-"
	}
	return string(result)
}

func (p *Position) Fen() string {
	var (
		color     = "w"
		enPassant = "-"
	)
	if p.NextPlayer == Black {
		color = "b"
	}
	if p.EnPassantTarget != 0 {
		enPassant = strings.ToLower(p.EnPassantTarget.String())
	}
	return fmt.Sprintf(
		"%s %s %s %s %d %d",
		p.Placement.Fen(),
		color,
		p.CastlingAvailability.Fen(),
		enPassant,
		p.HalfMoves,
		p.FullMoves+1,
	)
}

func ParseFenPlacement(str string) (Placement, error) {
	ranksInReverse := strings.SplitN(str, "/", 9)
	if len(ranksInReverse) != 8 {
		return nil, ErrInvalidRecord
	}
	p := make(Placement)
	for rank := range ranksInReverse {
		file := 0
		for _, char := range []rune(ranksInReverse[7-rank]) {
			if piece, ok := PiecesByLetter[char]; ok {
				p[NewCoordinate(file, rank)] = piece
				file++
			} else if char >= '1' && char <= '8' {
				file += int(char - '0')
			} else {
				return nil, ErrInvalidRecord
			}
		}
		if file != 8 {
			return nil, ErrInvalidRecord
		}
	}
	return p, nil
}

func MustParseFenPlacement(str string) Placement {
	p, err := ParseFenPlacement(str)
	if err != nil {
		panic(err)
	}
	return p
}

func ParseFenPosition(str string) (position *Position, err error) {
	fields := strings.SplitN(str, " ", 7)
	if len(fields) != 6 {
		return nil, ErrInvalidRecord
	}
	pos := new(Position)

	if pos.FullMoves, err = strconv.Atoi(fields[5]); err != nil {
		return
	}
	if pos.FullMoves < 1 {
		return nil, ErrInvalidRecord
	}
	pos.FullMoves--

	if pos.HalfMoves, err = strconv.Atoi(fields[4]); err != nil {
		return
	}
	if pos.HalfMoves < 0 {
		return nil, ErrInvalidRecord
	}

	if fields[3] != "-" {
		if pos.EnPassantTarget, err = ParseCoordinate(fields[3]); err != nil {
			return
		}
	}

	if pos.CastlingAvailability, err = ParseFenCastlingAvailability(fields[2]); err != nil {
		return
	}

	switch fields[1] {
	case "W", "w":
		pos.NextPlayer = White
	case "B", "b":
		pos.NextPlayer = Black
	default:
		return nil, ErrInvalidRecord
	}

	if pos.Placement, err = ParseFenPlacement(fields[0]); err != nil {
		return
	}

	position = pos
	return
}

func MustParseFenPosition(str string) *Position {
	p, err := ParseFenPosition(str)
	if err != nil {
		panic(err)
	}
	return p
}

func ParseFenCastlingAvailability(str string) (CastlingAvailability, error) {
	if str == "-" {
		return CastlingAvailability{}, nil
	}
	if str == "" || len(str) > 4 {
		return nil, ErrInvalidRecord
	}
	result := make(CastlingAvailability, len(str))
	for _, char := range []byte(str) {
		switch {
		case char == 'K':
			result[H1] = true
		case char == 'Q':
			result[A1] = true
		case char == 'k':
			result[H8] = true
		case char == 'q':
			result[A8] = true
		case char >= 'A' && char <= 'H':
			result[NewCoordinate(int(char-'A'), 0)] = true
		case char >= 'a' && char <= 'h':
			result[NewCoordinate(int(char-'a'), 7)] = true
		default:
			return nil, ErrInvalidRecord
		}
	}
	return result, nil
}

func MustParseFenCastlingAvailability(str string) CastlingAvailability {
	a, err := ParseFenCastlingAvailability(str)
	if err != nil {
		panic(err)
	}
	return a
}
