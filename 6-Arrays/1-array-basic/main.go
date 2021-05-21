package main

import (
	"fmt"
)

// the array declaration specifies what data type can be stored within array elements
// Go arrays are fixed size
// once created, cannot be add or remove elements 

func main() {

	// array declared using var keyword followed by 'arrayName' followed by [] brackets that hold `int` number of required elements followed by the 'dataType' the elements contain

	// var arrayName [ nElements ] dataType

	var vehicles [3]string

	// assign values to each array element
	vehicles[0] = "skateboard"
	vehicles[1] = "car"
	vehicles[2] = "bicycle"

	// array declared and intialized with values by stating values for each element in an 'array literal'
	// arrayName := [ nElements ] dataType { value, value, value, value }

	// 2-dimensional array:
	// var arrayName [ nElements ] [ nElements ] dataType

	// or, 2-dimensional array intialized and values assigned:
	// arrayName := [ nElements ] [ nElements ] dataType {{ value, value, }, { value, value, }}
	// individual elements value referenced by 'arrayName' followed by brackets representing each index:
	// arrayName[0][1]

	coordinates := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Vehicles:", vehicles)
	fmt.Println("Third Vehicle:", vehicles[2])
	fmt.Println("Coordinatess (X1,Y1):", coordinates[0][0])
	fmt.Println("Coordinatess (X2,Y3):", coordinatess[1][2])
}

//	% go run main.go
//	Vehicles: [skateboard car bicycle]
//	Third Vehicle: bicycle
//	Coordinatess (X1,Y1): 1
//	Coordinatess (X2,Y3): 6
