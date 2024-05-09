package colors

import (
	"fmt"
)

type Color string

const (
	Cyan   Color = "\033[38;5;159m"
	Pink   Color = "\033[38;5;225m"
	Red    Color = "\033[38;5;224m"
	Green  Color = "\033[38;5;194m"
	Yellow Color = "\033[38;5;230m"
	Purple Color = "\033[38;5;141m"
	Orange Color = "\033[38;5;216m"
	Reset  Color = "\033[0m"
)

const (
	BrightCyan   Color = "\033[1m\033[38;5;159m"
	BrightPink   Color = "\033[1m\033[38;5;225m"
	BrightRed    Color = "\033[1m\033[38;5;224m"
	BrightGreen  Color = "\033[1m\033[38;5;194m"
	BrightYellow Color = "\033[1m\033[38;5;230m"
	BrightPurple Color = "\033[1m\033[38;5;141m"
	BrightOrange Color = "\033[1m\033[38;5;216m"
)

func Format(messageColor Color, message ...any) string {
	text := fmt.Sprint(message...)

	return fmt.Sprintf("%v%v%v", messageColor, text, Reset)
}

// --> Color.FormatCyan("message")
func FormatCyan(message ...any) string {
	return Format(Cyan, message...)
}

func FormatPink(message ...any) string {
	return Format(Pink, message...)
}

func FormatRed(message ...any) string {
	return Format(Red, message...)
}

func FormatYellow(message ...any) string {
	return Format(Yellow, message...)
}
