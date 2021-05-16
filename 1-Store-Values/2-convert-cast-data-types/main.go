package main

import (
	"fmt"
	"strconv" // contains functions that convert to and from the string data type
)

// convert data types
// convert the value stored in a variable to a different data type

func main() {

	var str string = "42"
	var char byte = 'A'

	// "strconv.Atoi()" function converts a string into an "int" data type
	// returns the "int" equivalent of the string
	// Like many Other Go functions:
	// "strconv.Atoi()" and also returns a 2nd value that will be the "nil" zero value when conversion succeeds or error if fail
	// so, must assign the call to "strconv.Atoi()" func to 2 variables to receive the 2 returned values

	// casting conversion between different numeric data types is stating the required type followed by parentheses containing the value to be converted
	// var num int = 42
	// var decimal float64 = float64(num)
	num, err := strconv.Atoi(str)
	fmt.Printf("num:%v %T %v \n", num, num, err)

	decimal := float64(num)
	fmt.Printf("decimal: %.2f %T \n", decimal, decimal)

	fmt.Printf("char:%v %T %v \n", char, char, string(char))
}

//	% go run main.go
//	num:42 int <nil> 
//	decimal: 42.00 float64 
//	char:65 uint8 A
