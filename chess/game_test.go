package chess_test

import (
	"fmt"
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestGame_HalfMoveClock(t *testing.T) {
	game := NewGameWithPlacement(Placement{
		H8: Black.King(),
		H1: White.King(),
		A8: Black.Rook(),
		A1: White.Rook(),
		B2: White.Pawn(),
	})
	Equals(t, 0, game.HalfMoveClock())

	for _, move := range []struct {
		Move      string
		HalfMoves int
	}{
		{"b2b3", 0},
		{"a8a7", 1},
		{"b3b4", 0},
		{"a7a6", 1},
		{"a1a2", 2},
		{"a6a5", 3},
		{"a2a5", 0},
	} {
		game.MustMove(MustParseUCI(move.Move))
		Assert(t, move.HalfMoves == game.HalfMoveClock(),
			fmt.Sprintf("%s should result in %d half moves (got %d)", move.Move, move.HalfMoves, game.HalfMoveClock()))
	}
}
