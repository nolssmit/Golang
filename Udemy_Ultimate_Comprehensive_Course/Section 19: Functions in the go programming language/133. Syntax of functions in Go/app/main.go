package main

import "fmt"

func main() {
	// no params, no returns
	foo()

	// 1 param, no returns
	bar("Todd")


	// 1 param, 1 return	
	s := aloha("Todd")
	fmt.Println(s)

	// 2 params, 2 returns
	s1, n := dogYears("Todd", 40)
	fmt.Println(s1, n)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

// no params, no returns
func foo() {
	fmt.Println("I am from foo")
}

// 1 param, no returns
func bar(s string) {
	fmt.Println("My name is", s)
}

// 1 param, 1 return
func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

// 2 params, 2 returns
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}
