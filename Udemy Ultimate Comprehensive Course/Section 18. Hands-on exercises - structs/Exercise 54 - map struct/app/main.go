package main

import (
	"fmt"
)

type favFlavor []string

type person struct {
	first string
	last  string
	//	favIC []string
	favIC favFlavor
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		//	favIC: []string{"chocolate", "strawberry"},
		favIC: favFlavor{"chocolate", "strawberry"},
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		favIC: []string{"strawberry", "butter-pecan"},
	}

	//	 fmt.Printf("%#v\n%#v\n", p1, p2)

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(p1.favIC)
	// fmt.Println(p2.favIC)

	//	fmt.Printf("\n%v %v likes ", p1.first, p1.last)
	//	for _, v := range(p1.favIC) {
	//		fmt.Printf("%v, ", v)
	//	}

	//	fmt.Printf("\n%v %v likes ", p2.first, p2.last)
	//	for _, v := range(p2.favIC) {
	//		fmt.Printf("%v, ", v)
	//	}

	
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	
	for _, v := range m {
		fmt.Println(v)
		fmt.Println("-------------------------------")
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, "likes", v2)
		}
		fmt.Println("================================")
	}
}

/*
Take the code from the previous exercise, then store the VALUES of
type person in a map with the KEY of last name. Access each value in the map.
Print out the values, ranging over the slice.
*/
