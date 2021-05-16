package main

import (
	"fmt"
)

// function with parameters
// review/as expected

func square(num int) {

	fmt.Println("Received Copy:", num)

	num = num * num
	fmt.Println("Modified Copy:", num)
}

func main() {

	num := 5
	fmt.Println("Original:", num)

	square(num)
	fmt.Println("Original:", num)
}

//	% go run main.go
//	Original: 5
//	Received Copy: 5
//	Modified Copy: 25
//	Original: 5
