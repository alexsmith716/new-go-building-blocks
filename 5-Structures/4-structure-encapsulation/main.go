package main

import (
	"fmt"
	"new-go-building-blocks/utils"
)

// represent objects in programs -OOP
// Encapsulation, Inheritance and Polymorphism is OOP

// an OOP program uses methods to express the properties and operatons of a data structure
// a `var`, `func` or method that is not publically accessible is 'encapsulated'

// Encapsulation is data-hiding in a package that is imported into a program
// only `var`, `func` or methods that are capitalized are exported to the program
// lowercase `var`, `func` or methods are inaccessible in the program (already used in '4-Functions > 9-function-import-package')


func main() {

	// create instance of struct `LengthWidthHeight`
	var dimensions utils.CalcAreaVolume

	// call accessible method `SetSize()` to assign values to `CalcAreaVolume` struct
	dimensions.SetStructFields(9, 4, 12)

	fmt.Println("Calculated Area (length x width):", dimensions.MethodArea())
	fmt.Println("Calculated Volume (length x width x height):", dimensions.MethodVolume())
}

//	% go run main.go
//	Calculated Area (length x width): 36
//	Calculated Volume (length x width x height): 432
