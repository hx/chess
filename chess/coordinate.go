package chess

import (
	"strings"
)

type Coordinate uint8

const (
	A1, B1, C1, D1, E1, F1, G1, H1 Coordinate = iota << 3,
		1 | iota<<3,
		2 | iota<<3,
		3 | iota<<3,
		4 | iota<<3,
		5 | iota<<3,
		6 | iota<<3,
		7 | iota<<3
	A2, B2, C2, D2, E2, F2, G2, H2
	A3, B3, C3, D3, E3, F3, G3, H3
	A4, B4, C4, D4, E4, F4, G4, H4
	A5, B5, C5, D5, E5, F5, G5, H5
	A6, B6, C6, D6, E6, F6, G6, H6
	A7, B7, C7, D7, E7, F7, G7, H7
	A8, B8, C8, D8, E8, F8, G8, H8
)

func NewCoordinate(file, rank int) Coordinate {
	return Coordinate(rank<<3 | file)
}

func ParseCoordinate(algebraic string) (Coordinate, error) {
	if len(algebraic) != 2 {
		return 0, ErrInvalidCoordinate
	}
	b := []byte(strings.ToUpper(algebraic))
	return constrainCoordinate(
		int(b[0]-'A'),
		int(b[1]-'1'),
	)
}

func MustParseCoordinate(algebraic string) Coordinate {
	coord, err := ParseCoordinate(algebraic)
	if err != nil {
		panic(err)
	}
	return coord
}

func (c Coordinate) File() int {
	return int(c) & 0b111
}

func (c Coordinate) Rank() int {
	return int(c) >> 3
}

func (c Coordinate) String() string {
	return string([]byte{
		'A' + byte(c.File()),
		'1' + byte(c.Rank()),
	})
}

func (c Coordinate) VectorTo(other Coordinate) Vector {
	return Vector{
		File: other.File() - c.File(),
		Rank: other.Rank() - c.Rank(),
	}
}

func (c Coordinate) VectorFrom(other Coordinate) Vector {
	return other.VectorTo(c)
}

func (c Coordinate) VectorsTo(others ...Coordinate) []Vector {
	result := make([]Vector, len(others))
	for i, other := range others {
		result[i] = c.VectorTo(other)
	}
	return result
}

func (c Coordinate) Shift(vector Vector, times int) (Coordinate, error) {
	return constrainCoordinate(
		c.File()+vector.File*times,
		c.Rank()+vector.Rank*times,
	)
}

func (c Coordinate) To(to Coordinate) MoveRequest {
	return MoveRequest{From: c, To: to}
}

func constrainCoordinate(file, rank int) (Coordinate, error) {
	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return 0, ErrInvalidCoordinate
	}
	return NewCoordinate(file, rank), nil
}
