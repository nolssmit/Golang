package main

import "fmt"

type aEmployee map[string]int

func main() {
	//var employee = map[string]int{"Mark": 10, "Sandy": 20} 	//a
	//append Employee({"Mark": 10, "Sandy": 20})			//a
	//var employee = aEmployee{"Mark": 10, "Sandy": 20}		//b
	//employee := aEmployee{"Mark": 10, "Sandy": 20} 		//c
	//employee := aEmployee{}					//d1  Empty map
	//employee := make(map[string]int)				//d1  Empty map
	employee := make(aEmployee) //d1 Empty map
	employee["nols"] = 72       //d2
	employee["pat"] = 76        //d3
	delete(employee, "nols")    //d4

	fmt.Println(employee)
}