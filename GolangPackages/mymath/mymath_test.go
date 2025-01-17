package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
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
