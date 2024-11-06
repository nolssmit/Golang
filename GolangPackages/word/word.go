// package word counts words in a string
package word

import "strings"

// UseCount returns a map of words and their number of occurrences
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	// write the code for this func
     return len(strings.Fields(s))
}