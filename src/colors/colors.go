package colors

import (
	"fmt"
)

const (
	Cyan   = "\u001b[38;5;159m"
	Pink   = "\u001b[38;5;225m"
	Red    = "\u001b[38;5;224m"
	Green  = "\u001b[38;5;194m"
	Yellow = "\u001b[38;5;230m"
	Purple = "\u001b[38;5;141m"
	Reset  = "\u001b[0m"
)

const (
	BackgroundCyan   = "\u001b[48;5;159m"
	BackgroundPink   = "\u001b[48;5;225m"
	BackgroundRed    = "\u001b[48;5;224m"
	BackgroundGreen  = "\u001b[48;5;194m"
	BackgroundYellow = "\u001b[48;5;230m"
	BackgroundPurple = "\u001b[48;5;141m"
)

func Format(color string, message ...any) string {
	text := fmt.Sprint(message...)

	return fmt.Sprintf("%v%v%v", color, text, Reset)
}

// --> Color.FormatCyan("message")
func FormatCyan(message ...any) string {
	return Format(Cyan, message...)
}

func FormatPink(message ...any) string {
	return fmt.Sprintf("%v%v%v", Pink, message, Reset)
}
