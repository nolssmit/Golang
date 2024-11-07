package word

import (
	"fmt"
	"testing"

	"github.com/nolssmit/Golang/GolangPackages/quote"
)


func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

// func TestUseCount(t *testing.T) {
// 	m := UseCount("one two three")
//
// 	if m != 2 {
// 		t.Error("got", m, "want", 2)
// 	}
// }

// TestCount tests the Count function
func TestCount(t *testing.T) {
	n := Count("one two three")

	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func ExampleUseCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}


