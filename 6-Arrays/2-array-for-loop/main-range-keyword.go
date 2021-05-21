package main

import (
	"fmt"
)

// for-loop using "range" keyword
// "range" shortcuts iterating over array and returns each element index & value

func main() {

	arr := [...]int{59, 22, 7, 90, 123}

	fmt.Println("number of elements:", len(arr))

	for i, v := range arr {
		arr[i] = v / 1
		fmt.Printf("arr[i]: %v \n", arr[i])
	}

	for i, v := range arr {
		fmt.Printf("index: %v value: %v \n", i, v)
	}
}

//	% go run main-range-keyword.go
//	number of elements: 5
//	arr[i]: 59 
//	arr[i]: 22 
//	arr[i]: 7 
//	arr[i]: 90 
//	arr[i]: 123 
//	index: 0 value: 59 
//	index: 1 value: 22 
//	index: 2 value: 7 
//	index: 3 value: 90 
//	index: 4 value: 123
