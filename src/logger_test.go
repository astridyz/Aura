package aura

import "testing"

func TestLoggerPrint(t *testing.T) {
	log := NewLogger("test")

	log.Print("Hey, keep going! Don't give up.")
}
