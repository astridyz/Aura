package aura

import (
	"fmt"
	"os"
	"strings"

	"github.com/astridyz/Aura/src/colors"
)

type logger struct {
	Prefix *Prefix
}

// Formats prefix with color, adds it to a subprefix and adds a message
func (l *logger) log(s *subPrefix, color colors.Color, message ...any) string {

	var prefix string
	var subPrefixText string

	if l.Prefix != nil {
		prefix = l.Prefix.Format()
	}

	if s != nil {
		subPrefixText = colors.Format(s.Color, s.Structure)
	}

	text := colors.Format(color, message...)
	result := removeSpaceFromFormatedMessage(fmt.Sprintf("%v%v %v\n", prefix, subPrefixText, text))

	return result
}

// --> Print

// Is equivalent to print() with color and prefix if defined.
func (l *logger) Print(message ...any) string {
	result := l.log(nil, colors.Cyan, message...)

	print(result)
	return result
}

// Is equivalent to logger.Print() with formatting
func (l *logger) Printf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return l.Print(text)
}

// --> Warn

// Is equivalent to print() with color and prefix if defined.
func (l *logger) Warn(message ...any) string {
	subPrefix := &subPrefix{
		Structure: "WARN ",
		Color:     colors.BrightYellow,
	}
	result := l.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	return result
}

// Is equivalent to logger.Print() with formatting
func (l *logger) Warnf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return l.Warn(text)
}

// --> Panic

// It prints a message and call panic() with color
func (l *logger) Panic(message ...any) {
	subPrefix := &subPrefix{
		Structure: "PANIC",
		Color:     colors.BrightPurple,
	}
	result := l.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	panic(fmt.Sprint(message...))
}

// Equivalent to logger.Panic() with formatting
func (l *logger) Panicf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Panic(text)
}

// --> Error

// Prints an error message with color and prefix
func (l *logger) Error(message ...any) {
	subPrefix := &subPrefix{
		Structure: "ERROR",
		Color:     colors.BrightRed,
	}
	result := l.log(subPrefix, colors.BrightYellow, message...)

	print(result)
}

// Equivalent to logger.Error() with formatting
func (l *logger) Errorf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Error(text)
}

// --> Fatal

// Prints an error message with color and prefix and calls os.Exit(1)
func (l *logger) Fatal(message ...any) {
	subPrefix := &subPrefix{
		Structure: "FATAL",
		Color:     colors.BrightOrange,
	}
	result := l.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	os.Exit(1)
}

// Equivalent to logger.Fatal() with formatting
func (l *logger) Fatalf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Fatal(text)
}

// --> Setters

// Sets the logger prefix.
// The prefix will be added to every logger function message.
func (l *logger) SetPrefix(Prefix *Prefix) {
	l.Prefix = Prefix
}

// --> Prefix

// It's added to every logger function message
type Prefix struct {
	Structure string
	Color     colors.Color
}

// Formats prefix with color
func (p *Prefix) Format() string {
	return colors.Format(p.Color, p.Structure)
}

// Some functions have subprefixes, like error and panic
// It is used to symbolize which message type was printed
type subPrefix struct {
	Structure string
	Color     colors.Color
}

// --> log := aura.NewLogger()

// Creates a new logger
func NewLogger(prefixes ...*Prefix) *logger {
	instance := &logger{}

	if len(prefixes) > 0 {
		instance.Prefix = prefixes[0]
	}

	return instance
}

// --> Utils

func removeSpaceFromFormatedMessage(message string) string {
	message = strings.TrimLeftFunc(message, func(r rune) bool {
		return r == ' '
	})

	return message
}
