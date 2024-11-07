package word

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/quote"
	"testing"
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

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

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