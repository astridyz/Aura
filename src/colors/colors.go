package colors

import (
	"fmt"
)

type Color string

const (
	Cyan   Color = "\u001b[38;5;159m"
	Pink   Color = "\u001b[38;5;225m"
	Red    Color = "\u001b[38;5;224m"
	Green  Color = "\u001b[38;5;194m"
	Yellow Color = "\u001b[38;5;230m"
	Purple Color = "\u001b[38;5;141m"
	Orange Color = "\u001b[38;5;216m"
	Reset  Color = "\u001b[0m"
)

const (
	BrightCyan   Color = "\u001b[1m\u001b[38;5;159m"
	BrightPink   Color = "\u001b[1m\u001b[38;5;225m"
	BrightRed    Color = "\u001b[1m\u001b[38;5;224m"
	BrightGreen  Color = "\u001b[1m\u001b[38;5;194m"
	BrightYellow Color = "\u001b[1m\u001b[38;5;230m"
	BrightPurple Color = "\u001b[1m\u001b[38;5;141m"
	BrightOrange Color = "\u001b[1m\u001b[38;5;216m"
)

func Format(mensageColor Color, message ...any) string {
	text := fmt.Sprint(message...)

	return fmt.Sprintf("%v%v%v", mensageColor, text, Reset)
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
