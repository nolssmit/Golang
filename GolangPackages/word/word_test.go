package word

import (
	"fmt"
//	"strings"
	"testing"
)

const myQuote = `Robert Cohn was once once middleweight boxing champion Robert`

// func TestUseCount(myQuote string) map[string]int {
// 	xs := strings.Fields(myQuote)

// 	m := make(map[string] int)
// 	for _, v := range xs {
// 		m[v]++
// 	}

// 	if m !=  map[Cohn:1 Robert:1 boxing:1 champion:1 middleweight:1 once:1 was:1] 	 
// 	{
// 				t.Error("got", m, "want", xs)
// 	}		
// }	


// TestCount tests the Count function
func TestCount(t *testing.T) {	
	n := Count(myQuote)
    fmt.Println("myQuote: ", myQuote) 

	if n != 9 {
		t.Error("got", n, "want", 9)
	}
}