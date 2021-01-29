package chess

type Color bool

const (
	Black Color = false
	White Color = true
)

func (c Color) String() string {
	if c {
		return "White"
	}
	return "Black"
}

func (c Color) New(pieceType PieceType) Piece {
	return NewPiece(c, pieceType)
}

func (c Color) Pawn() Piece   { return c.New(Pawn) }
func (c Color) Knight() Piece { return c.New(Knight) }
func (c Color) Bishop() Piece { return c.New(Bishop) }
func (c Color) Rook() Piece   { return c.New(Rook) }
func (c Color) Queen() Piece  { return c.New(Queen) }
func (c Color) King() Piece   { return c.New(King) }
