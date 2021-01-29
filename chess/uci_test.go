package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestMoveRequest_UCI(t *testing.T) {
	Equals(t, "D2D4", D2.To(D4).UCI())
	Equals(t, "H7H8Q", H7.To(H8).As(Queen).UCI())
}

func TestParseUCI(t *testing.T) {
	Equals(t, D2.To(D4), MustParseUCI("D2d4"))
	Equals(t, H7.To(H8).As(Queen), MustParseUCI("h7H8q"))
}
