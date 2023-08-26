package helpers

import "log"

func ErrorCheck(err error) {  // Note Uppercase to access it from outside this file
	if err != nil {
		log.Fatal(err)
	}
}