// Package saying provides a function to greet people
package saying

import (
	"fmt"
	"testing"
)

// Greet returns a greeting to the user
func Greet(s string) string {	
	return fmt.Sprint("Welcome my dear ", s)
}

func TestGreet(t *testing.T) {	
	s := Greet("Nols Smit")

	if s != "Welcome my dear Nols Smit" {
		t.Error("got", s, "expected","Welcome my dear Nols Smit")
	}
}

func ExampleGreet() {	
	fmt.Println(Greet("Nols Smit"))
	// Output: Welcome my dear Nols Smit
}

func BenchmarkGreet(b *testing.B) {	
	for i := 0; i < b.N; i++ {
		Greet("Nols Smit")
	}
}