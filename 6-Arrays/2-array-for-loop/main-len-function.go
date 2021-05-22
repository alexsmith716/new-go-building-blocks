package main

import (
	"fmt"
)
 
// for-loop using Go's built-in `len()` function
// `len()` function accepts array as argument and returns an `int` that is the number of elements in array
// `len()` function will exit the loop when it reaches the array's final element
 

func main() {

	var arr [6]int

	fmt.Println("number of elements:", len(arr))

	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("index: %v value: %v \n", i, arr[i])
	}
}

//	% go run main-len-function.go
//	number of elements: 6
//	index: 0 value: 0 
//	index: 1 value: 1 
//	index: 2 value: 4 
//	index: 3 value: 9 
//	index: 4 value: 16 
//	index: 5 value: 25
