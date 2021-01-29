package chess_test

import (
	. "github.com/hx/chess/chess"
	. "github.com/hx/chess/chess/testing"
	"testing"
)

func TestParseCoordinate(t *testing.T) {
	coord, err := ParseCoordinate("A1")
	Equals(t, A1, coord)
	Ok(t, err)

	coord, err = ParseCoordinate("b4")
	Equals(t, B4, coord)
	Ok(t, err)

	coord, err = ParseCoordinate("H8")
	Equals(t, H8, coord)
	Ok(t, err)

	coord, err = ParseCoordinate("")
	Equals(t, A1, coord)
	Equals(t, ErrInvalidCoordinate, err)

	coord, err = ParseCoordinate("A9")
	Equals(t, A1, coord)
	Equals(t, ErrInvalidCoordinate, err)

	coord, err = ParseCoordinate("foobar")
	Equals(t, A1, coord)
	Equals(t, ErrInvalidCoordinate, err)
}

func TestCoordinate_String(t *testing.T) {
	Equals(t, "A1", A1.String())
	Equals(t, "F5", F5.String())
	Equals(t, "H8", H8.String())
}
