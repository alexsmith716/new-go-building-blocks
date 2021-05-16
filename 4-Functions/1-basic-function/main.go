package main

import (
	"fmt"
)

// basic function


// variables declared inside a func or block are accessible within that block - they are local scoped

func first() {
	msg := "Hello from the 1st function!"
	fmt.Println(msg) // var `msg` only accessible inside func `first()`
}

func sqFive() {
	fmt.Printf("%v \n", 5*5) // (format specifier, variable)
}

// once func statements have been executed, program flow resumes at the point following the func call
func main() {
	first()
	fmt.Print("5 x 5 = ")
	sqFive()
}

//	% go run main.go
//	Hello from the 1st function!
//	5 x 5 = 25
