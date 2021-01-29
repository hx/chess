package chess

type Piece struct {
	Color Color
	Type  PieceType
}

func NewPiece(color Color, pieceType PieceType) Piece {
	return Piece{
		Color: color,
		Type:  pieceType,
	}
}

func (p *Piece) String() string {
	return p.Color.String() + " " + p.Type.String()
}

// Uppercase for white, lowercase for black
func (p *Piece) Letter() (letter rune) {
	letter = p.Type.Letter()
	if p.Color == Black {
		letter += 'a' - 'A'
	}
	return
}

func (p *Piece) Unicode() rune {
	return p.Type.Figurine(p.Color)
}

func (p *Piece) Is(color Color, pieceType PieceType) bool {
	return p != nil && p.Type == pieceType && p.Color == color
}

func ParsePieceTypeLetter(letter rune) (PieceType, error) {
	if piece, ok := PiecesByLetter[letter]; ok {
		return piece.Type, nil
	}
	return 0, ErrInvalidPieceType
}

func MustParsePieceTypeLetter(letter rune) PieceType {
	t, err := ParsePieceTypeLetter(letter)
	if err != nil {
		panic(err)
	}
	return t
}
