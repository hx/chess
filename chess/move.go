package chess

import (
	"strings"
)

type Move struct {
	MoveRequest
	*Piece
	CapturedPiece *Piece
	IsEnPassant   bool
	IsCheck       bool
}

func (m *Move) IsCastling() bool {
	return m.Piece.Type == King && m.From.File() == 4 && (m.To.File() == 2 || m.To.File() == 6)
}

func (m *Move) IsCapture() bool {
	return m.CapturedPiece != nil
}

func (m *Move) CapturedPieceCoordinate() Coordinate {
	if !m.IsCapture() {
		return 0
	}
	if m.IsEnPassant {
		return NewCoordinate(m.To.File(), m.From.Rank())
	}
	return m.To
}

func (m *Move) IsPromotion() bool {
	return m.PromoteTo != 0
}

func (m *Move) CastlingMoveRequest() (req MoveRequest) {
	if !m.IsCastling() {
		panic("not a castling move")
	}
	var rank = m.From.Rank()
	if m.To.File() == 2 {
		req.From = NewCoordinate(0, rank)
		req.To = NewCoordinate(3, rank)
	} else {
		req.From = NewCoordinate(7, rank)
		req.To = NewCoordinate(5, rank)
	}
	return
}

func (m *Move) String() string {
	str := new(strings.Builder)
	str.WriteString(m.Piece.Color.String())
	str.WriteRune(' ')
	if m.IsPromotion() {
		str.WriteString(Pawn.String())
	} else {
		str.WriteString(m.Piece.Type.String())
	}
	str.WriteRune(' ')
	str.WriteString(m.From.String())
	str.WriteString(" to ")
	str.WriteString(m.To.String())
	if m.IsCapture() {
		str.WriteString(", capture ")
		str.WriteString(m.CapturedPiece.Type.String())
		if m.IsEnPassant {
			str.WriteString(" en passant")
		}
	}
	if m.IsPromotion() {
		str.WriteString(", promote to ")
		str.WriteString(m.PromoteTo.String())
	}
	if m.IsCastling() {
		str.WriteString(", ")
		if m.To.File() == 2 {
			str.WriteString("queen")
		} else {
			str.WriteString("king")
		}
		str.WriteString("side castle")
	}
	if m.IsCheck {
		str.WriteString(", check")
	}
	return str.String()
}
