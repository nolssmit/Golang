package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make: "Toyota",
		model: "Corolla",
		doors: 4,
		color: "White",

	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make: "Ford",
		model: "Mustang",
		doors: 2,
		color: "Red",

	}

	fmt.Println(v1.model, v1.engine.electric)
	fmt.Println(v1)
	fmt.Println("----------------------------------")
	fmt.Println(v2.model, v2.engine)
	fmt.Println(v2)
}

/*
Create a type engine struct, and include this field
electric bool

Create a type vehicle struct, and include these fields
engine
make
model
doors
color

Create two VALUES of TYPE vehicle
use a composite literal

Print out each of these values.
Print out a single field from each of these values.
*/