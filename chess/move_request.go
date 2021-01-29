package chess

type MoveRequest struct {
	From      Coordinate
	To        Coordinate
	PromoteTo PieceType
}

func (r MoveRequest) As(promoteTo PieceType) MoveRequest {
	r.PromoteTo = promoteTo
	return r
}
