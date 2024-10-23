// Package dog contains functions for working with dogs.
package dog

import (
	"strings"
)

// WhenGrownUp returns the passed in string in uppercase.
func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
// Years returns the passed in number of years in dog years.
func Years(y int) int {
	return  y * 7
}	