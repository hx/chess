package chess

var Language string

var i18n = map[string]map[PieceType]struct {
	Letter rune
	Name   string
}{
	"en": {
		Pawn:   {'P', "Pawn"},
		Knight: {'N', "Knight"},
		Bishop: {'B', "Bishop"},
		Rook:   {'R', "Rook"},
		Queen:  {'Q', "Queen"},
		King:   {'K', "King"},
	},
}

var PiecesByLetter map[rune]Piece

func SetLanguage(language string) error {
	if _, ok := i18n[language]; !ok {
		return ErrUnknownLanguage
	}
	Language = language
	PiecesByLetter = make(map[rune]Piece, 12)
	for _, pieceType := range []PieceType{Pawn, Knight, Bishop, Rook, Queen, King} {
		for _, color := range []Color{Black, White} {
			piece := color.New(pieceType)
			PiecesByLetter[piece.Letter()] = piece
		}
	}
	return nil
}

func MustSetLanguage(language string) {
	if err := SetLanguage(language); err != nil {
		panic(err)
	}
}

func init() {
	MustSetLanguage("en")
}
