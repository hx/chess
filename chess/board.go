package chess

type Board [64]*Square

func (b *Board) At(coordinate Coordinate) *Square {
	return b[coordinate]
}

func NewEmptyBoard() *Board {
	return Placement{}.Board()
}

func (b *Board) Placement() Placement {
	p := make(Placement)
	for _, square := range b {
		if square.IsOccupied() {
			p[square.Coordinate] = *square.Piece
		}
	}
	return p
}
