package main

import "testing"

func TestMysum(t *testing.T) {
	x := mySum(1, 2, 3)
	if x != 6 {
		t.Error("Expected 6, got", x)
	}
	if mySum(1, 2, 3, 4, 5, 6, 7) != 28 {
		t.Error("Expected 28, got", mySum(1, 2, 3, 4, 5, 6, 7))
	}
	if mySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) != 55 {
		t.Error("Expected 55, got", mySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}
}