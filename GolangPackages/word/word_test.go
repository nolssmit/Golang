package word

import (
//	"fmt"
//	"strings"
	"testing"
)

const myQoute = `Robert Cohn was once once middleweight boxing champion Robert`

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

func TestCount(t *testing.T) {	
	n := Count(myQoute)

	if n != 70 {
		t.Error("got", n, "want", 7)
	}
}