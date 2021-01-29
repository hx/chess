package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestSquare_BackgroundColor(t *testing.T) {
	Equals(t, Black, NewSquare("A1").BackgroundColor())
	Equals(t, Black, NewSquare("B2").BackgroundColor())
	Equals(t, White, NewSquare("B1").BackgroundColor())
	Equals(t, White, NewSquare("A2").BackgroundColor())
	Equals(t, Black, NewSquare("H8").BackgroundColor())
}
