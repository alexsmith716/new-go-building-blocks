package main

import (
	"fmt"
	"math"
)

// 'math' functions for rounding a decimal

// functions:
// math.Floor(x), math.Ceil(x), math.Round(x), math.Trunc(x)

// https://golang.org/pkg/builtin/#float64
// float64 is the set of all IEEE-754 64-bit floating-point numbers.
// https://floating-point-gui.de/formats/fp/

// "in Go, a name is exported (from a package) if it begins with a capital letter -"Pi" & "E")"
// "when importing a package, you can refer only to its exported names" `math.Pi` & `math.E`

// cast a type 'float64' to type 'integer': 
// Go does not support implicit type casting
// type conversion bewtwen compatible data types can be performed
// syntax 'dataType(value)' to convert a value to another data type

func main() {

	// intialize floating point variable `pi`
	var pi float64 = math.Pi
	fmt.Println("math.Pi:", pi)

	fmt.Println("math.Floor(pi):", math.Floor(pi))
	fmt.Println("math.Ceil(pi):", math.Ceil(pi))
	fmt.Println("math.Round(pi):", math.Round(pi))
	fmt.Println("math.Trunc(pi):", math.Trunc(pi))

	// reduce decimal places of floating-point value `pi`
	pi = pi*100
	pi = math.Round(pi)
	pi = pi/100
	// reducedPi = math.Round(pi*100)/100)
	fmt.Println("reduce decimal places of 'math.Pi':", pi)

	var exponential1 float64 = math.E
	// var exponential1 float64 = math.Exp(1)
	fmt.Printf("math.Exp(1): %v %T \n", exponential1, exponential1)

	// cast 'float64' `exponential1` to an 'integer' data type
	// converting a floating-point to an integer will truncate and remove decimal values
 	var exponential2 int = int(exponential1)
	fmt.Printf("cast 2.718281828459045 float64 to interger: %v %T \n", exponential2, exponential2)
}

//	% go run main.go
//	math.Pi: 3.141592653589793
//	math.Floor(pi): 3
//	math.Ceil(pi): 4
//	math.Round(pi): 3
//	math.Trunc(pi): 3
//	reduce decimal places of 'math.Pi': 3.14
//	math.Exp(1): 2.718281828459045 float64 
//	cast 2.718281828459045 float64 to interger: 2 int
