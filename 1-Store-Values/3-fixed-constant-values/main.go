package main

import (
	"fmt"
)

// fix constant values
// fixed value that never changes during execution of program
// using "const" not "var"
// a const declaration must be initialized in the declaration

// a constant declaration can use "const generator" "iota" to create a sequence of related const values
// first const gets assigned the "iota" value

func main() {

	const pi = 3.14159

	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)

	fmt.Printf("Pi approximately: %v \n", pi)
	fmt.Printf("Red: %v point \n", red)
	fmt.Printf("Blue: %v points \n", blue)
	fmt.Printf("Black: %v points \n", black)
}

//	% go run main.go
//	Pi approximately: 3.14159 
//	Red: 1 point 
//	Blue: 5 points 
//	Black: 7 points
