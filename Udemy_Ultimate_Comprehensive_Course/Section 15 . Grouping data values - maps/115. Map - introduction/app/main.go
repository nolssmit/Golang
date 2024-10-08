package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
		"Nols":72,
	}

	fmt.Println("The age of Henry was", am["Henry"])
	fmt.Printf("Nols's is %v jaar oud\n", am["Nols"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	
	an := make(map[string]int)

	// aj := make(map[string]int)
	// aj["Nols"] = 72
	aj := map[string]int {"Nols": 72}	
	fmt.Printf("aj %#v\n",aj)
	fmt.Printf("Length of map aj: %v\n", len(aj``))
	// a slice sou wees: aj := make([]string, <starting length>, <capacity>)
	// aj[0] = "Nols Smit"
	// of, aj := [] string {"Nols Smit"}
	
	an["Lucas"] = 28
	an["Steph"] = 37
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	fmt.Println(len(an))



	/*
		m := map[string]int{
			"todd":  42,
			"henry": 16,
		}

		//  --- or ---

		// m := make(map[string]int)
		// m["todd"] = 42
		// m["henry"] = 16

		fmt.Println("Henry's age is ", m["henry"])

		for k := range m {
			fmt.Printf("just the keys: %s\n", k)
		}

		for k, v := range m {
			fmt.Printf("%s is %d years old\n", k, v)
		}

		for _, v := range m {
			fmt.Printf("just the values: %d\n", v)
		}

		if v, ok := m["Padget"]; ok {
			fmt.Printf("%s is %d years old\n", "Padget", v)
		} else {
			fmt.Printf("%s not found\n", "Padget")
		}

		//delete
		m["Shakespeare"] = 459
		fmt.Printf("%#v\n", m)
		delete(m, "Shakespeare")
		fmt.Printf("%#v\n", m)
		delete(m, "Shakespeare") // no panic

		// len
		fmt.Println("len: ", len(m))
	*/
}

//  --- or ---

/*
var m map[string]int
fmt.Println(m["tunde"])
m = make(map[string]int)
m["todd"] = 42
m["henry"] = 16
*/
