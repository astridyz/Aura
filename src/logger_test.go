package aura

import "testing"

func TestPrinting(t *testing.T) {
	log := New()
	log.Print("Testing something very cool")
}
