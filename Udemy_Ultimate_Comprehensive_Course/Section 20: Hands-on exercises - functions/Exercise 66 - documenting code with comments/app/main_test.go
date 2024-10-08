package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := Paradise("Bali")
	want := "My idea of paradise is Bali"
	if got != want {
		log.Fatalf("Error, got: %s and want %s", got, want)
	}
}

// See: https://chat.openai.com/c/ea38ff4b-3bba-4148-b42f-6f35dc5a404d
