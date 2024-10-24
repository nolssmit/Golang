package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age`
}

func main(){
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	fmt.Println("=======================================")

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("All of the data:", people)

	fmt.Println("=======================================")

	fmt.Printf("%+v\n", people)

	fmt.Println("=======================================")

	for i, v := range people{
		fmt.Println("Person number:", i)
		fmt.Println(v.First, v.Last, v.Age)
		fmt.Println("------------------")
	}
}