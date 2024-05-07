package colors

import "fmt"

const (
	cyan   = "\u001b[38;5;159m"
	pink   = "\u001b[38;5;225m"
	red    = "\u001b[38;5;224m"
	green  = "\u001b[38;5;194m"
	yellow = "\u001b[38;5;230m"
	purple = "\u001b[38;5;141m"
	reset  = "\u001b[0m"
)

/*
const (
	backgroundCyan   = "\u001b[48;5;159m"
	backgroundPink   = "\u001b[48;5;225m"
	backgroundRed    = "\u001b[48;5;224m"
	backgroundGreen  = "\u001b[48;5;194m"
	backgroundYellow = "\u001b[48;5;230m"
	backgroundPurple = "\u001b[48;5;141m"
)
*/

// --> Color.Cyan("m")

func FormatCyan(message string) string {
	return fmt.Sprintf("%v%v%v\n", cyan, message, reset)
}
