package main_test

import (
	"testing"
	"github.com/nolssmit/Golang/GolangPackages/acdc"
)

// TestSum let you run a test on the Sum function in this package 
func TestSum(t *testing.T) {
	s := acdc.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if s != 55 {
		t.Errorf("Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) = %d; want 55", s)
	}
}
