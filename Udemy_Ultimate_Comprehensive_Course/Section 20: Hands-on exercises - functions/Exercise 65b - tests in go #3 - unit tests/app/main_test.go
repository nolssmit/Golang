package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := paradise("Bali")
	want := "My idea of paradise is Bali"
	if got != want {
		log.Fatalf("Error, got: %s and want %s", got, want)
	}
}
