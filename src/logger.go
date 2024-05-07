package aura

import (
	"fmt"

	"github.com/astridyz/aura/colors"
)

type logger struct {
	Name   string
	Prefix *Prefix
}

func (l *logger) Print(message ...any) string {
	prefix := l.Prefix.Format()
	text := colors.FormatCyan(message...)
	result := fmt.Sprintf("%v%v\n", prefix, text)

	print(result)
	return result
}

func (l *logger) Printf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return l.Print(text)
}

func (l *logger) Panic(message ...any) {
	prefix := l.Prefix.Format()
	text := colors.FormatYellow(message...)

	errorText := colors.Format(colors.BrightPurple, "PANIC")

	result := fmt.Sprintf("%v%v %v\n", prefix, errorText, text)

	print(result)
	panic(fmt.Sprint(message...))
}

func (l *logger) Panicf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Panic(text)
}

func (l *logger) Error(message ...any) {
	prefix := l.Prefix.Format()
	text := colors.FormatYellow(message...)
	errorText := colors.Format(colors.BrightRed, "ERROR")

	result := fmt.Sprintf("%v%v %v\n", prefix, errorText, text)

	print(result)
}

func (l *logger) Errorf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	l.Error(text)
}

// --> Setters

func (l *logger) SetPrefix(Prefix *Prefix) {
	l.Prefix = Prefix
}

type Prefix struct {
	Structure string
	Color     colors.Color
}

func (p *Prefix) Format() string {
	return colors.Format(p.Color, p.Structure)
}

// --> log := aura.NewLogger()
func NewLogger(name string) *logger {
	instance := &logger{
		Name: name,
	}

	return instance
}
