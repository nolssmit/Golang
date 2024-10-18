package puppy

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From13(){
	fmt.Println("I'm from version 1.3.0")
}