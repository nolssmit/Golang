package main

import "fmt"

func intDelta(n *int){
	*n = 43
}

func sliceDelta(ii []int){
	ii[0] = 99
}

func mapDelta(md map[string]int, s string){
	md[s] = 72
}

func main(){
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["Nols"] = 73
	fmt.Println(m["Nols"])
	mapDelta(m, "Nols")
	fmt.Println(m["Nols"])
}