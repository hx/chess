package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestBoard_Fen(t *testing.T) {
	Equals(t, "8/8/8/8/8/8/8/8", NewEmptyBoard().Fen())
	Equals(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", NewGame().Board.Fen())
}

func TestCastlingAvailability_Fen(t *testing.T) {
	Equals(t, "KQkq", CastlingAvailability{A1: true, H1: true, A8: true, H8: true}.Fen())
	Equals(t, "KQq", CastlingAvailability{A1: true, H1: true, A8: true, H8: false}.Fen())
	Equals(t, "Kq", CastlingAvailability{A1: false, H1: true, A8: true, H8: false}.Fen())
	Equals(t, "DEad", CastlingAvailability{D1: true, E1: true, A8: true, D8: true}.Fen())
	Equals(t, "BEef", CastlingAvailability{B1: true, E1: true, E8: true, F8: true}.Fen())
	Equals(t, "-", CastlingAvailability{}.Fen())
}

func TestPosition_Fen(t *testing.T) {
	game := NewGame()
	for _, move := range []struct {
		Move string
		Fen  string
	}{
		// Taken from https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
		{"", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"e2e4", "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"},
		{"c7c5", "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"},
		{"g1f3", "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"},
	} {
		if move.Move != "" {
			game.MustMove(MustParseUCI(move.Move))
		}
		Equals(t, move.Fen, game.Position().Fen())
	}
}

func TestParseFenCastlingAvailability(t *testing.T) {
	for _, str := range []string{"KQkq", "KQq", "KQk", "KQ", "Kq", "Kkq", "Qkq", "k", "-", "DEad", "BEef", "Agh"} {
		Equals(t, str, MustParseFenCastlingAvailability(str).Fen())
	}
}

func TestParseFenPosition(t *testing.T) {
	for _, str := range []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
	} {
		Equals(t, str, MustParseFenPosition(str).Fen())
	}
}
