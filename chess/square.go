package chess

type Square struct {
	Coordinate
	Piece *Piece
}

func NewSquare(algebraicCoordinate string) *Square {
	return &Square{Coordinate: MustParseCoordinate(algebraicCoordinate)}
}

func (s *Square) IsVacant() bool {
	return s.Piece == nil
}

func (s *Square) IsOccupied() bool {
	return s.Piece != nil
}

func (s *Square) BackgroundColor() Color {
	return (s.Rank()+s.File())%2 == 1
}
