package aura

import "fmt"

type logger struct {
	Prefix *Prefix
}

// --> Prefix is used before log messages,
// --> It has a structure and a color to define how it'll be showed
type Prefix struct {
	Body  string
	color string
}

//
type message struct {
	Body  any
	Level logLevel
}

// --> Color type, it needs to be formated to be used in strings
type color = int

// --> Used to send a color to formats, without needing to know the color integer
type TextColor = string

// --> Log level type
type logLevel = int

// --> Formated colors are colors that were formated after
// --> using the colorIntToString function()
type formatedColor = string

var (
	// Colors map[string]color

	formatedColors map[logLevel]string

	// brightFormatedColors map[logLevel]string
)

// --> Colors constants, all of them are pastel colors
const (
	Cyan   color = 159
	Pink   color = 225
	Red    color = 224
	Yellow color = 230
	Purple color = 141
	Orange color = 216
)

const ColorReset string = "\033[0m"

// --> Log levels constants, defining the color of the messages
const (
	Critical logLevel = iota + 1
	Error
	Warning
	Debug
	Info
)

// --> Initializating module
func init() {
	initColors()
}

// --> Instanciating a new logger and returning it
// --> Factory function
func New(prefix ...*Prefix) *logger {
	instance := &logger{}

	if len(prefix) > 0 {
		instance.Prefix = prefix[0]
	}

	return instance
}

func (l *logger) log(msg message) {
	text := formatColor(msg)
	print(text)
}

func (l *logger) Print(args ...any) {
	l.log(message{Body: args, Level: Info})
}

// --> Initing all colors in a map
func initColors() {
	formatedColors = map[logLevel]string{
		Critical: colorIntToString(Orange),
		Error:    colorIntToString(Red),
		Warning:  colorIntToString(Yellow),
		Debug:    colorIntToString(Pink),
		Info:     colorIntToString(Cyan),
	}

	/*
		brightFormatedColors = map[logLevel]string{
			Critical: colorIntToBrightString(Orange),
			Error:    colorIntToBrightString(Red),
			Warning:  colorIntToBrightString(Yellow),
			Debug:    colorIntToBrightString(Pink),
			Info:     colorIntToBrightString(Cyan),
		}

		Colors = map[TextColor]color{
			"Cyan":   Cyan,
			"Pink":   Pink,
			"Red":    Red,
			"Yellow": Yellow,
			"Purple": Purple,
			"Orange": Orange,
		}
	*/
}

func formatColor(msg message) string {
	text := fmt.Sprint(msg.Body)
	formatedColor := formatedColors[msg.Level]
	print(formatedColor)

	return fmt.Sprintf("%v%v%v\n", formatedColor, text, ColorReset)
}

// --> Used to format colors integers into ANSI codes
func colorIntToString(color color) formatedColor {
	return fmt.Sprintf("\033[38;5;%vm", color)
}

/*
// --> Same as colorIntToString(), but it adds a bright ANSI Code,
// --> turning the text BOLD.
func colorIntToBrightString(color color) formatedColor {
	return fmt.Sprintf("\033[1m\033[38;5;%vm", color)
}
*/
