package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestFigurines(t *testing.T) {
	// See https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode

	Equals(t, rune(0x2654), Figurines[White][King])
	Equals(t, rune(0x2655), Figurines[White][Queen])
	Equals(t, rune(0x2656), Figurines[White][Rook])
	Equals(t, rune(0x2657), Figurines[White][Bishop])
	Equals(t, rune(0x2658), Figurines[White][Knight])
	Equals(t, rune(0x2659), Figurines[White][Pawn])

	Equals(t, rune(0x265a), Figurines[Black][King])
	Equals(t, rune(0x265b), Figurines[Black][Queen])
	Equals(t, rune(0x265c), Figurines[Black][Rook])
	Equals(t, rune(0x265d), Figurines[Black][Bishop])
	Equals(t, rune(0x265e), Figurines[Black][Knight])
	Equals(t, rune(0x265f), Figurines[Black][Pawn])
}
