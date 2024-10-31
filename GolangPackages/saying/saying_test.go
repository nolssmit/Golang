package saying_test

import (
	"fmt"
	"testing"

	"github.com/nolssmit/Golang/GolangPackages/saying"
)

func TestGreet(t *testing.T) {	
	s := saying.Greet("Nols Smit")

	if s != "Welcome my dear Nols Smit" {
		t.Error("got", s, "expected","Welcome my dear Nols Smit")
	}
}

func ExampleGreet() {	
	fmt.Println(saying.Greet("Nols Smit"))
	// Output: Welcome my dear Nols Smit
}

func BenchmarkGreet(b *testing.B) {	
	for i := 0; i < b.N; i++ {
		saying.Greet("Nols Smit")
	}
}