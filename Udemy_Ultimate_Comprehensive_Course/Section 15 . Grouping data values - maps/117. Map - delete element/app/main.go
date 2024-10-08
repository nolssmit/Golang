package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
		"Nols":   72,
	}

	fmt.Println("The age of Henry was", am["Henry"])
	fmt.Printf("Nols's is %v jaar oud\n", am["Nols"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	// delete an element from a map
	delete(am, "Nols")
	fmt.Printf("%#v\n", am)

	fmt.Println("------- accessing keys that don't exist --------")
	fmt.Println(am["XXX"])
	fmt.Println("------------------------------------------------")
	fmt.Println("-------- handling not existing map keys --------")
	v, ok := am["xxx"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key: xxx doesn't exist in map am")
	}
	fmt.Println("-----------------------------------------------")

	an := make(map[string]int)

	// aj := make(map[string]int)
	// aj["Nols"] = 72
	aj := map[string]int{"Nols": 72}
	fmt.Printf("aj %#v\n", aj)
	fmt.Printf("Length of map aj: %v\n", len(aj))
	// a slice sou wees: aj := make([]string, <starting length>, <capacity>)
	// aj[0] = "Nols Smit"
	// of, aj := [] string {"Nols Smit"}
	aj["Pat"] = 76
	fmt.Printf("Map aj: %#v\n", aj)
	// for range over a map
	for k, v := range aj {
		fmt.Printf("Key: %v\tValue: %v\n", k, v)
	}

	for k := range aj {
		fmt.Printf("Key: %v\n", k)
	}

	//for range over a slice
	xi := []int{42, 43, 44, 45, 46, 47}
	fmt.Printf("Slice: %#v\n", xi)
	for i := 0; i < len(xi); i++ {
		fmt.Println(xi[i])
	}
	for i, v := range xi {
		fmt.Printf("Key: %v\tValue: %v\n", i, v)
	}

	for i := range xi {
		fmt.Printf("Key: %v\n", i)
	}

	fmt.Println("-------- Delete element 44 from the slice ---------")
	yi := append(xi[:2], xi[3:]...)
	fmt.Printf("%#v\n", yi)
	fmt.Println("---------------------------------------------------")

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
