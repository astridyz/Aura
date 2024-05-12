package aura

import "testing"

func TestPrinting(t *testing.T) {
	log := New(&Prefix{Body: "Astrid:", Level: Debug})

	log.Print("Testing something very cool")

	log.Warn("Im warning you!")

	log.Error("Yayyyy that's an error!")

	log.Panic("RIP. Panic.")
}
