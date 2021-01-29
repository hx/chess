package chess

import (
	"bytes"
	"github.com/hx/chess/chess/ansi"
	"io"
)

type BoardColorScheme struct {
	Piece      map[Color]ansi.Color
	Background ansi.Color
}

type TextColorScheme struct {
	Text       ansi.Color
	Background ansi.Color
}

type BoardFormatter struct {
	BoardScheme map[Color]BoardColorScheme
	RankScheme  TextColorScheme
	FileScheme  TextColorScheme
}

func NewBoardFormatter() *BoardFormatter {
	return &BoardFormatter{
		BoardScheme: map[Color]BoardColorScheme{
			White: {
				Piece: map[Color]ansi.Color{
					White: ansi.BrightWhite,
					Black: ansi.Black,
				},
				Background: ansi.White,
			},
			Black: {
				Piece: map[Color]ansi.Color{
					White: ansi.BrightWhite,
					Black: ansi.Black,
				},
				Background: ansi.BrightBlack,
			},
		},
		RankScheme: TextColorScheme{
			Text:       ansi.White,
			Background: ansi.Black,
		},
		FileScheme: TextColorScheme{
			Text:       ansi.White,
			Background: ansi.Black,
		},
	}
}

func (f *BoardFormatter) write(board *Board) []byte {
	w := new(bytes.Buffer)
	for rank := 7; rank >= 0; rank-- {
		w.WriteString(ansi.ColorString(f.RankScheme.Text, f.RankScheme.Background))
		w.WriteRune(' ')
		w.WriteRune('1' + rune(rank))
		w.WriteRune(' ')
		for file := 0; file <= 7; file++ {
			square := board.At(NewCoordinate(file, rank))
			var pieceColor Color
			if square.IsOccupied() {
				pieceColor = square.Piece.Color
			}
			boardColor := f.BoardScheme[square.BackgroundColor()]
			w.WriteString(ansi.ColorString(boardColor.Piece[pieceColor], boardColor.Background))
			if square.IsOccupied() {
				w.WriteRune(' ')
				w.WriteRune(square.Piece.Unicode())
				w.WriteRune(' ')
			} else {
				w.WriteString("   ")
			}
		}
		w.WriteString(ansi.Reset)
		w.WriteRune('\n')
	}
	w.WriteString(ansi.ColorString(f.FileScheme.Text, f.FileScheme.Background))
	w.WriteString("    a  b  c  d  e  f  g  h ")
	w.WriteString(ansi.Reset)
	w.WriteRune('\n')
	return w.Bytes()
}

func (f *BoardFormatter) WriteTo(board *Board, writer io.Writer) (int, error) {
	return writer.Write(f.write(board))
}

func (f *BoardFormatter) String(board *Board) string {
	return string(f.write(board))
}
