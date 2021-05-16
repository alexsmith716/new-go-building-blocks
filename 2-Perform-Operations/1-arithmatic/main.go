package main

import (
	"fmt"
)

// operators assignment, addition, subtraction, multiplication and division act expected
// the "%" (modulus/modulo) remainder operator divides the first given number by the second and returns the remainder
// all operators same as in JS

func main() {
	a := 8
	b := 4

	// "\t" escape sequence is used to output a tab space to align results
	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Remainder:\t", (a % b))

	a++
	fmt.Println("Increment:\t", a)
	b--
	fmt.Println("Decrement:\t", b)
}

//	% go run main.go
//	Addition:	 12
//	Subtraction:	 4
//	Multiplication:	 32
//	Division:	 2
//	Remainder:	 0
//	Increment:	 9
//	Decrement:	 3
