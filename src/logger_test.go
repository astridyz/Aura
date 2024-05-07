package aura

import "testing"

func TestLoggerPrint(t *testing.T) {
	log := NewLogger("test")

	log.Print("test")
}
