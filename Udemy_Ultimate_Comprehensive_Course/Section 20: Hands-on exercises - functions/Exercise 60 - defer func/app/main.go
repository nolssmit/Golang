package main

import (
    "fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main() {
    fmt.Println("Printing from main") 
    defer print1()
    defer print2()
    defer print3()
    defer print4()
}

func print1() {
    fmt.Println("Print 1")
}

func print2() {
    fmt.Println("Print 2")
}

func print3() {
    fmt.Println("Print 3")
}

func print4() {
    fmt.Println("Print 4")
}

/*
  defered functions run in LIFO order
  Last in First out
*/
