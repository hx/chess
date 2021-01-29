package chess

type CastlingAvailability map[Coordinate]bool

func NewCastlingAvailability(placement Placement) CastlingAvailability {
	a := make(CastlingAvailability)
	for coord, piece := range placement {
		if piece.Type == Rook {
			a[coord] = true
		}
	}
	return a
}

func (a CastlingAvailability) Copy() CastlingAvailability {
	b := make(CastlingAvailability, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func (a CastlingAvailability) Unset(coordinate Coordinate) {
	delete(a, coordinate)
}
