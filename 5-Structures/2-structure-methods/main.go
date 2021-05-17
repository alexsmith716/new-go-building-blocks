package main

import (
	"fmt"
)

// functions can be attached to structs
// a function attached to a struct is a method of the struct
// functions are attached by adding a 'receiver' to the function signiture
// the receiver specifies a parameter name and the struct name as its type

// func (paramName structName) funcName()

// the method can be a 'value' receiver OR a 'pointer' receiver
// a value receiver affects the copy of the struct
// a pointer receiver affects the original struct fields

// func (paramName *structName) funcName()

// refer to '1-Store-Values' > '4-point-to-stored-values'

// objects are created in Go by defining a struct (the struct is a template)
// copies/instances are then made from that struct/template

// struct instances are modified with attribute values/fields and behaviors/methods

type coffee struct {
	color string
	flavor  string
}

// attach method `experience()` to struct `coffee`
func (c coffee) experience() string {
	return "amazing"
}

func main() {

	darkRoast := coffee{color: "dark", flavor: "rich"}
	colombian := coffee{color: "light", flavor: "mellow"}

	fmt.Println("Dark Roast coffee beans are", darkRoast.color)
	fmt.Println("Dark Roast coffee flavor is", darkRoast.flavor)
	// call method 
	fmt.Println("Dark Roast coffee is", darkRoast.experience())

	fmt.Println("Colombian coffee beans are", colombian.color)
	fmt.Println("Colombian coffee flavor is", colombian.flavor)
	fmt.Println("Colombian coffee is", colombian.experience())
}

//	% go run main.go
//	Dark Roast coffee beans are dark
//	Dark Roast coffee flavor is rich
//	Dark Roast coffee is amazing
//	Colombian coffee beans are light
//	Colombian coffee flavor is mellow
//	Colombian coffee is amazing
