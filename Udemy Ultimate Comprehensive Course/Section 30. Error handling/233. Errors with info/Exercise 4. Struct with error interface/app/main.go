// Package main is the main package
package main

import (
	"fmt"
	"log"
)

// nogateMathError is an error type
type nogateMathError struct {
	lat  string
	long string	
	err error
}

func (n nogateMathError) Error() string {
	return fmt.Sprintf("a norgate error occured: %v %v %v",n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, nogateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library
//
// https://golang.org/pkg/net/#OpError
// https://golang.org/src/pkg/net/dial.go
// https://golang.org/src/pkg/net/net.go
//
// https://golang.org/src/pkg/encoding/json/decode.go