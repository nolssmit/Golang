package dog 

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {	
	x := Years(7)

	if x != 70 {
		t.Error("got", x, "want", 70)
	}
}

func TestYearsTwo(t *testing.T) {	
	x := YearsTwo(7)

	if x != 140 {
		t.Error("got", x, "want", 140)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 140
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}	
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}	
}