package aura

import (
	"fmt"

	"github.com/astridyz/aura/colors"
)

type logger struct {
	Name   string
	Prefix Prefix
}

func (l *logger) Print(message ...any) {
	prefix := colors.Format(l.Prefix.Color, l.Prefix.Structure)
	text := colors.FormatCyan(message...)

	print(fmt.Sprintf("%v %v\n", prefix, text))
}

/*
func (l *logger) Start() {
}
*/

type Prefix struct {
	Structure string
	Color     string
}

// --> log := aura.NewLogger()
func NewLogger(name string) *logger {
	instance := &logger{
		Name: name,
		Prefix: Prefix{
			Structure: "Astrid:",
			Color:     colors.BackgroundPink,
		},
	}

	return instance
}

func FormatPrefix(Prefix Prefix) {

}
