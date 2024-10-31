package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {	
	s :=Greet("Nols Smit")

	if s != "Welcome my dear Nols Smit" {
		t.Error("got", s, "expected","Welcome my dear Nols Smit")
	}
}

func ExampleGreet() {	
	fmt.Println(Greet("Nols Smit"))
	// Output: 
	// Welcome my dear Nols Smit
}

func BenchmarkGreet(b *testing.B) {	
	for i := 0; i < b.N; i++ {
		Greet("Nols Smit")
	}
}