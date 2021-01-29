package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestGame_EnPassant(t *testing.T) {
	game := NewGameWithPlacement(MustParseFenPlacement("7k/8/8/8/1p6/8/P7/7K"))
	_, err := game.Move(A2.To(A4))
	Ok(t, err)

	// Or:
	//game := NewGameWithPosition(MustParseFenPosition("7k/8/8/8/Pp6/8/8/7K b - a3 0 1"))

	move, err := game.Move(B4.To(A3))
	Ok(t, err)
	Assert(t, move.IsEnPassant, "Move should be en passant")
	Assert(t, move.CapturedPiece.Is(White, Pawn), "White pawn should be captured")
}

func TestGame_IllegalEnPassant(t *testing.T) {
	game := NewGameWithPlacement(MustParseFenPlacement("7k/8/8/8/1p6/P7/8/7K"))
	_, err := game.Move(A3.To(A4))
	Ok(t, err)
	move, err := game.Move(B4.To(A3))
	Equals(t, ErrMoveNotAllowed, err)
	Assert(t, move == nil, "Move should not happen")
}

func TestGame_KingsideCastle(t *testing.T) {
	game := NewGameWithPlacement(Placement{
		E8: Black.King(),
		E1: White.King(),
		H1: White.Rook(),
	})
	move, err := game.Move(E1.To(G1))
	Ok(t, err)
	Assert(t, move.IsCastling(), "Move should be castling")
	Assert(t, game.At(F1).Piece.Is(White, Rook), "Rook should have moved to F1")
	Assert(t, game.At(H1).IsVacant(), "H1 should be vacant")
}

func TestGame_QueensideCastle(t *testing.T) {
	game := NewGameWithPlacement(Placement{
		E8: Black.King(),
		E1: White.King(),
		A1: White.Rook(),
	})
	move, err := game.Move(E1.To(C1))
	Ok(t, err)
	Assert(t, move.IsCastling(), "Move should be castling")
	Assert(t, game.At(D1).Piece.Is(White, Rook), "Rook should have moved to D1")
	Assert(t, game.At(A1).IsVacant(), "A1 should be vacant")
}
