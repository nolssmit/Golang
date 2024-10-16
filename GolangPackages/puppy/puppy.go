package puppy

import (
	"github.com/nolssmit/Golang/GolangPackages/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrounUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrounUp(Barks())
}