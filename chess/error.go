package chess

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrMoveNotAllowed     Error = "Move not allowed"
	ErrInvalidCoordinate  Error = "Invalid coordinate"
	ErrNoFirstMove        Error = "No first move"
	ErrGameHasEnded       Error = "Game has ended"
	ErrInvalidMoveRequest Error = "Invalid move request"
	ErrInvalidPieceType   Error = "Invalid piece type"
	ErrInvalidRecord      Error = "Invalid record"
	ErrUnknownLanguage    Error = "Unknown language"
)
