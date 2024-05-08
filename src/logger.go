package aura

import (
	"fmt"

	"github.com/astridyz/Aura/src/colors"
)

type logger struct {
	Name   string
	prefix *Prefix
}

// --> Print

// Is equivalent to print() with color and prefix if defined.
func (l *logger) Print(message ...any) string {
	result := l.prefix.FormatText(message...)

	print(result)
	return result
}

// Is equivalent to logger.Print() with formatting
func (l *logger) Printf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return l.Print(text)
}

// --> Panic

// It prints a message and call panic() with color
func (l *logger) Panic(message ...any) {
	subPrefix := &subPrefix{
		Structure: "PANIC",
		Color:     colors.BrightPurple,
	}
	result := l.prefix.FormatTextWithSubPrefix(subPrefix, message...)

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
	result := l.prefix.FormatTextWithSubPrefix(subPrefix, message...)

	print(result)
}

// Equivalent to logger.Error() with formatting
func (l *logger) Errorf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Error(text)
}

// --> Setters

// Sets the logger prefix.
// The prefix will be added to every logger function message.
func (l *logger) SetPrefix(Prefix *Prefix) {
	l.prefix = Prefix
}

// --> Prefix

// It's added to every logger function message
type Prefix struct {
	Structure string
	Color     colors.Color
}

// Some functions have subprefixes, like error and panic
// It is used to symbolize which message type was printed
type subPrefix struct {
	Structure string
	Color     colors.Color
}

// Formats prefix with color
func (p *Prefix) Format() string {
	return colors.Format(p.Color, p.Structure)
}

// Formats prefix with color and adds it to a message
func (p *Prefix) FormatText(message ...any) string {
	prefix := p.Format()
	text := colors.FormatCyan(message...)
	result := fmt.Sprintf("%v%v\n", prefix, text)

	return result
}

// Formats prefix with color, adds it to a subprefix and adds a message
func (p *Prefix) FormatTextWithSubPrefix(sub *subPrefix, message ...any) string {
	prefix := p.Format()
	text := colors.FormatYellow(message...)
	subPrefixText := colors.Format(sub.Color, sub.Structure)

	result := fmt.Sprintf("%v%v %v\n", prefix, subPrefixText, text)

	return result
}

// --> log := aura.NewLogger()

// Creates a new logger
func NewLogger(name string) *logger {
	instance := &logger{
		Name: name,
	}

	return instance
}
