package main

import (
	"fmt"
)

// operator precedence

func main() {

	sum := 100*3 + 4 - 5
	fmt.Printf("Default Order: %v \n", sum)

	sum = 2 * ((3 + 4) - 5)
	fmt.Printf("Forced Order1: %v \n", sum)

	sum = 7 % 3 * 2
	fmt.Printf("Default Order2: %v \n", sum)

	sum = 7 % (3 * 2)
	fmt.Printf("Forced Order3: %v \n", sum)
}

//	nothing new here
//	force compiler to ignore default precedence levels by using parentheses
//	precedence (add parentheses to mofify operation precedence):
//	1) * / % << >> &
//	2) + - | ^
//	3) == != < <= >=
//	4) &&
//	5) ||

//	% go run main.go
//	Default Order: 299 
//	Forced Order: 4 
//	
//	Default Order: 2 
//	Forced Order: 1
