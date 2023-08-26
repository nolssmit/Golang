package main

import (
	"fmt"
)

// ----- STRUCTS -----
type customer struct {
	name    string
	address string
	bal     float64
}

// This struct has a function associated
type rectangle struct{
	length, height float64	
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owns us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, addres string) {
	c.address = addres
}

// Struct composition : Putting a struct in another
type contact struct {
	fname string
	lname string
	phone string
}
	
type business struct {
	name string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s  %s\n",
			b.name, b.contact.fname, b.contact.phone)	
}



func main()  {
	// ----- STRUCTS -----
	// Structs allow you to store values with many data types

	// Add values
	var tS customer
	tS.name = "Nols Smit"
	tS.address = "5 Main st"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// ... or as reference
	newCustAdd(&tS, "123 South St")
	fmt.Println("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main rd", 0.0}
	fmt.Println("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	fmt.Println("Rect Area :",rect1.Area())

	// Go doesn't support inheritance, but it does support
	// composition by embedding a struct in another
	con1 := contact {
		"James", "Wang", "555-1212",
	}
	bus1 := business{
		"ABC Plumbing",
		"234 Nort St",
		con1,
	}
	bus1.info()
}