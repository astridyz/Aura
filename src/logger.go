package aura

import (
	"fmt"
	"os"
	"strings"
)

// --> Main logger
// --> Main interface to logging
type logger struct {
	Prefix *Prefix
}

// --> Prefix is used before log messages,
// --> It has a structure and a color to define how it'll be showed
type Prefix struct {
	Body  string
	Level logLevel
}

// --> Formats prefix in a colored text with its prefix
func (p *Prefix) format() string {
	return formatColor(p.Level, true, p.Body)
}

// --> Color type, it needs to be formated to be used in strings
type color = int

// --> Log level type
type logLevel = int

// --> Formated colors are colors that were formated after
// --> using the colorIntToString function()
type formatedColor = string

var (
	formatedColors map[logLevel]string

	brightFormatedColors map[logLevel]string
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

// --> Final ANSI Code to reset string color back to default
const ColorReset string = "\033[0m"

// --> Log levels constants, defining the color of the messages
const (
	Critical logLevel = iota + 1
	Panic
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

// --> Main function to logging
// --> It defines the message that will be printed
func (l *logger) log(s *Prefix, level logLevel, args ...any) {

	var formatedLoggerPrefix string
	var result string
	var formatedText = formatColor(level, false, args...)

	if l.Prefix != nil {
		formatedLoggerPrefix = l.Prefix.format()
	}

	if s != nil {
		result = removeSpaces(fmt.Sprintf("%v %v %v\n", formatedLoggerPrefix, s.format(), formatedText))
	} else {
		result = removeSpaces(fmt.Sprintf("%v %v\n", formatedLoggerPrefix, formatedText))
	}
	print(result)
}

// --> Prints a message in the terminal
func (l *logger) Print(args ...any) {
	l.log(nil, Info, args...)
}

// --> Same as l.Print() but with formatting pattern
func (l *logger) Printf(format string, args ...any) {
	l.log(nil, Info, fmt.Sprintf(format, args...))
}

// --> Prints an warn message.
// --> Doesn't stop code from running.
// --> Useful when calling sensitive functions
func (l *logger) Warn(args ...any) {
	l.log(&Prefix{Body: "WARN", Level: Warning}, Warning, args...)
}

// --> Same as l.Warn() but with formatting pattern
func (l *logger) Warnf(format string, args ...any) {
	l.log(&Prefix{Body: "WARN", Level: Warning}, Warning, fmt.Sprintf(format, args...))
}

// --> Prints an error message.
// --> Doesn't stop code from running.
func (l *logger) Error(args ...any) {
	l.log(&Prefix{Body: "ERROR", Level: Error}, Debug, args...)
}

// --> Same as l.Error() but with formatting pattern
func (l *logger) Errorf(format string, args ...any) {
	l.log(&Prefix{Body: "ERROR", Level: Error}, Debug, fmt.Sprintf(format, args...))
}

// --> Prints an error message.
// --> Call panic()
func (l *logger) Panic(args ...any) {
	l.log(&Prefix{Body: "PANIC", Level: Panic}, Warning, args...)
	panic(fmt.Sprint(args...))
}

// --> Same as l.Panic() but with formatting pattern
func (l *logger) Panicf(format string, args ...any) {
	l.log(&Prefix{Body: "PANIC", Level: Panic}, Warning, fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

// --> Prints an error message.
// --> Call os.exit(1)
func (l *logger) Fatal(args ...any) {
	l.log(&Prefix{Body: "FATAL", Level: Critical}, Error, args...)
	os.Exit(1)
}

// --> Same as l.Panic() but with formatting pattern
func (l *logger) Fatalf(format string, args ...any) {
	l.log(&Prefix{Body: "FATAL", Level: Critical}, Error, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// --> Initing all colors in a map
func initColors() {
	formatedColors = map[logLevel]string{
		Critical: colorIntToString(Orange),
		Panic:    colorIntToString(Purple),
		Error:    colorIntToString(Red),
		Warning:  colorIntToString(Yellow),
		Debug:    colorIntToString(Pink),
		Info:     colorIntToString(Cyan),
	}

	brightFormatedColors = map[logLevel]string{
		Critical: colorIntToBrightString(Orange),
		Panic:    colorIntToBrightString(Purple),
		Error:    colorIntToBrightString(Red),
		Warning:  colorIntToBrightString(Yellow),
		Debug:    colorIntToBrightString(Pink),
		Info:     colorIntToBrightString(Cyan),
	}
}

// --> Formats a text into a colored text
func formatColor(level logLevel, bright bool, args ...any) string {
	text := fmt.Sprint(args...)

	var formatedColor formatedColor
	if bright {
		formatedColor = brightFormatedColors[level]
	} else {
		formatedColor = formatedColors[level]
	}

	result := fmt.Sprintf("%v%v%v", formatedColor, text, ColorReset)
	return result
}

// --> Remove spaces from beginning messages
// --> It occurs when there's no prefix or subprefix
// --> So that's function is a must when getting results
func removeSpaces(message string) string {
	message = strings.TrimLeftFunc(message, func(r rune) bool {
		return r == ' '
	})

	return message
}

// --> Used to format colors integers into ANSI codes
func colorIntToString(color color) formatedColor {
	return fmt.Sprintf("\033[38;5;%vm", color)
}

// --> Same as colorIntToString(), but it adds a bright ANSI Code,
// --> turning the text BOLD.
func colorIntToBrightString(color color) formatedColor {
	return fmt.Sprintf("\033[1m\033[38;5;%vm", color)
}
