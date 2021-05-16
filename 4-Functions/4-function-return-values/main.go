package main

import (
	"fmt"
)

// return values

//	function basic syntax

//	func funcName(paramName type) returnType {
//		statements-to-be-executed
//		return returnValue
//	}

// Go functions can return multiple values
// data type of each value must be specified as comma separated list in parentheses after the parameter list
// the values must be listed in the same order

//	func funcName(paramName type) (returnType, returnType) {
//		statements-to-be-executed
//		return returnValue, returnValue
//	}

// assign any unwanted returned values to "_" underscore character (blank identifier)

func main() {

	_, b, c := cube(5)
	// a, b, c := cube(5)
	// only display "b" & "c"
	fmt.Println(b, "Cubed =", c)
	// fmt.Println(a, b, "Cubed =", c)
}

func cube(num int) (string, int, int) {
	return "Result", num, (num * num * num)
}

//	% go run main.go
//	5 Cubed = 125
