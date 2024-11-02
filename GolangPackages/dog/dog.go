// Package dog contains functions for working with dogs.
package dog

import (
	"strings"
)

// WhenGrownUp returns the passed string in uppercase.
func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
// Years returns the passed in number of years in dog years
// where one human year is 10 dog years.
func Years(y int) int {
	return  y * 10
}	

// YearsTwo returns the passed in number of years in dog years
// where one human year is 20 dog years.
func YearsTwo(y int) int {
	return  y * 20
}	