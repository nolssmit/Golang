package main

import (
	"fmt"
)

type Pear interface {
	allTheFruits()
}

type Orange struct {
	valueInsideOfOrange int
}

func (o *Orange) allTheFruits() {
	fmt.Printf("Apples value : %d\r\n", o.valueInsideOfOrange)
}

type Apple struct {
	valueInsideOfApple int
}

func (a *Apple) allTheFruits() {
	fmt.Printf("Apples value : %d\r\n", a.valueInsideOfApple)
}

func (a *Apple) appleOnlyFunction() {
	fmt.Printf("I can not be used inside a pear.\r\n")
}

func ICanTakeValuesThatLookLikePears(p Pear) {
	p.allTheFruits()
	//p.appleOnlyFunction()
}

func ICanOnlyTakeApples(a *Apple) {
	a.appleOnlyFunction()
}

func ChangeValueViaPointer(a *Apple) {
	// a will point to what was passed in
	a.valueInsideOfApple = 2
}

func ChangeValueNotUsingPointer(a Apple) {
	// a is a copy of what was passed in, therefore a is a shiny brand new apple, and has had it's values copied from the passed in version
	a.valueInsideOfApple = 10
}

func main() {
	a := &Apple{1}
	o := &Orange{valueInsideOfOrange: 100}

	// can execute directly the function associated to Apple
	a.allTheFruits()

	// I can pass apple to a pear interface as it complies with its interface (allTheFruits)
	ICanTakeValuesThatLookLikePears(a)
	// I can also pass orange to pear interface too.
	ICanTakeValuesThatLookLikePears(o)

	// this is a apple only function
	ICanOnlyTakeApples(a)

	// cannot pass orange into a apple only function
	//ICanOnlyTakeApples(o)

	// Non-Pointer examples
	aa := Apple{10}
	aa.allTheFruits()
	// aa will be copied, and passed into the function, hence, changes are not done on the object in this scope
	ChangeValueNotUsingPointer(aa)
	aa.allTheFruits()

	// Pointer example
	aa.valueInsideOfApple = 50
	aa.allTheFruits()

	// Now we will pass aa's address, so the pointer in the function will "point" to this object, and will alter its value
	ChangeValueViaPointer(&aa)
	aa.allTheFruits()
}