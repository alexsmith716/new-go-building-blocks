package main

import (
	"fmt"
)

// for-loop
// review/as expected

// "\t" escape sequence is used to output a tab space to align results
func main() {

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop Iteration", counter)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("\nOuter Loop Iteration", i)

		for j := 1; j <= 3; j++ {
			fmt.Println("\tInner Loop Iteration", j)
		}
	}
}

//	% go run main.go
//	Loop Iteration 1
//	Loop Iteration 2
//	Loop Iteration 3
//	Loop Iteration 4
//	Loop Iteration 5
//	
//	Outer Loop Iteration 1
//		Inner Loop Iteration 1
//		Inner Loop Iteration 2
//		Inner Loop Iteration 3
//	
//	Outer Loop Iteration 2
//		Inner Loop Iteration 1
//		Inner Loop Iteration 2
//		Inner Loop Iteration 3
//	
//	Outer Loop Iteration 3
//		Inner Loop Iteration 1
//		Inner Loop Iteration 2
//		Inner Loop Iteration 3
