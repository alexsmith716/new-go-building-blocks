package main

import (
	"fmt"
)

// anonymous function & closures (function literal/lambda)
// review/as expected

// a function definition that does not contain a name
// used as closures (a nested anonymous function that retains access to (private) vars declared in outer function)

func main() {

	// assign anonymous func to `area`
	area := func(length, width int) int {
		return length * width
	}

	fmt.Printf("area Type: %T \n", area)
	fmt.Println("Area 1:", area(10, 4))
	fmt.Println("Area 2:", area(12, 5))

	// anonymous IIFE assigned to `counter`
	// return a function which is returning int
	// 'higher-order' function (returning the anonymous function)
	counter := func() func() int {
		num := 0
		return func() int {
			num++
			return num
		}
	}()

	fmt.Printf("counter type: %T \n", counter)

	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())
}

//	% go run main.go
//	area Type: func(int, int) int 
//	Area 1: 40
//	Area 2: 60

//	counter type: func() int 
//	Count: 1
//	Count: 2
//	Count: 3
