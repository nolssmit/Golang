package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int{4, 8, 15, 16, 23, 42}
	v := CenteredAvg(xi)

	if v != 15.5 {
		t.Error("got", v, "want 15.5")
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{4, 8, 15, 16, 23, 42})
	}
}

func ExampleCenteredAvg() {
	xi := []int{4, 8, 15, 16, 23, 42}
	v := CenteredAvg(xi)

	fmt.Println(v)
	// Output: 
	// 15.5
}	

