package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

	var vName string
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.5
	
	v4 = 2.6
	v4 = 2.7

	pl(vName, v1, v2, v3, v4)

	pl(reflect.TypeOf(25))

	pa := reflect.TypeOf(25)
	fmt.Println(pa)
	pl(reflect.TypeOf(25))


	cV1 := 1.6
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))
}