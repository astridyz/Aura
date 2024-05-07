package aura

import (
	"testing"

	"github.com/astridyz/aura/colors"
)

func TestLoggerPrint(t *testing.T) {

	log := NewLogger("test")
	log.SetPrefix(&Prefix{Structure: "Astrid: ", Color: colors.BrightGreen})

	log.Print("Hey, keep going! Don't give up.")
	log.Printf("Hey, %v, how are you?", "Kame")
}
