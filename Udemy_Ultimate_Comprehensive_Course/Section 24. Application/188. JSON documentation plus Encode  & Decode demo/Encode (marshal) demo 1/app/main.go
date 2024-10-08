package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "MonePenny",
		Age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("====================================")
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====================================")
	fmt.Println(string(bs))
	fmt.Println("====================================")
	os.Stdout.Write(bs)
}
