package chess

type Placement map[Coordinate]Piece

func (p Placement) Board() *Board {
	board := Board{}
	for i := range board {
		square := Square{Coordinate: Coordinate(i)}
		if piece, ok := p[square.Coordinate]; ok {
			p := piece
			square.Piece = &p
		}
		board[i] = &square
	}
	return &board
}
