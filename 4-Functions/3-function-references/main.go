package main

import (
	"fmt"
)

// passing references
// review values, address & pointers

// `square()` requires a int pointer `*ptrNum`
func square(ptrNum *int) {

	fmt.Println("Received Address ptrNum:", ptrNum)

	// modify original's memory address
	*ptrNum = *ptrNum * *ptrNum
	fmt.Println("Modified Original Value ptrNum:", *ptrNum)
	fmt.Printf("Modified Address ptrNum: %v \n", &ptrNum)
}

// declare & initialize a variable `num`
func main() {

	num := 5
	fmt.Println("Original Value num:", num)

	fmt.Printf("Original Address num: %v \n", &num)

	// pass memory address var `&num` (modify the original argument)
	// pass reference to the original, not a copy of its value
	square(&num)
	fmt.Println("Original Value num:", num)
}

//	% go run main.go
//	Original Value num: 5
//	Original Address num: 0xc000014070 
//	Received Address ptrNum: 0xc000014070
//	Modified Original Value ptrNum: 25
//	Modified Address ptrNum: 0xc00000e020 
//	Original Value num: 25
