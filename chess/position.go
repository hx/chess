package chess

type Position struct {
	Placement            Placement
	NextPlayer           Color
	CastlingAvailability CastlingAvailability
	EnPassantTarget      Coordinate
	HalfMoves            int

	// Starts at 0, one less than FEN
	FullMoves int
}

func NewPosition(placement Placement, nextPlayer Color) *Position {
	return &Position{
		Placement:            placement,
		NextPlayer:           nextPlayer,
		CastlingAvailability: NewCastlingAvailability(placement),
	}
}

const StartingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
