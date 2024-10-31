// Package saying provides a function to greet people
package saying

import (
	"fmt"
)

// Greet returns a greeting to the user
func Greet(s string) string {	
	return fmt.Sprint("Welcome my dear ", s)
}