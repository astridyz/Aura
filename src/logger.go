package aura

import "github.com/astridyz/aura/colors"

type logger struct {
	Name string
}

func (l *logger) Print(message string) {
	print(colors.FormatCyan(message))
}

func (l *logger) Start() {
	// --> TODO
}

// --> log := aura.NewLogger()
func NewLogger(name string) *logger {
	instance := &logger{
		Name: name,
	}

	return instance
}
