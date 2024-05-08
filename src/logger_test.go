package aura

import (
	"testing"
)

func TestLoggerError(t *testing.T) {

	log := NewLogger()

	/*
		log.SetPrefix(&Prefix{
			Structure: "Astrid: ",
			Color:     colors.BrightPink,
		})
	*/

	log.Print("Hey, keep going! Don't give up.")
	log.Printf("Hey, %v, how are you?", "Kame")

	log.Warn("Cmon.. dont be afraid!")
	log.Warnf("Hey, %v, be allert!", "Kame")

	log.Error("That's an error!")
	log.Errorf("Hey, %v, help me with this problem...", "Kame")

	log.Fatal("And that's a fatal.")
}
