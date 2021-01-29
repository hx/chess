package ansi

import (
	"fmt"
)

const Reset = "\033[0m"

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BrightBlack Color = iota + 60
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

func ColorString(foreground, background Color) string {
	return fmt.Sprintf("\033[%d;%dm", foreground+30, background+40)
}
