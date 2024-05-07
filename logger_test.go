package aura

import (
	"testing"

	"github.com/astridyz/Aura/colors"
)

func TestLoggerError(t *testing.T) {
	log := NewLogger("Astrid")

	log.SetPrefix(&Prefix{
		Structure: "Astrid: ",
		Color:     colors.BrightPink,
	})

	log.Print("Hey, keep going! Don't give up.")
	log.Printf("Hey, %v, how are you?", "Kame")

	log.Error("That's an error!")
	log.Errorf("Hey, %v, help me with this problem...", "Kame")

	log.Panic("And that's a panic.")
}
