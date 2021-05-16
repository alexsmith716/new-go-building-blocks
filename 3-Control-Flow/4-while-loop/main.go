package main

import (
	"fmt"
)

// while-loop
// review/as expected

func main() {

	counter := 1

	for counter <= 5 {
		fmt.Println("Loop Iteration", counter)
		counter++
	}

	i := 10
	for {
		fmt.Println("Countdown", i)
		i--
		if i < 1 {
			fmt.Println("Lift Off!")
			break
		}
	}
}

//	% go run main.go
//	Loop Iteration 1
//	Loop Iteration 2
//	Loop Iteration 3
//	Loop Iteration 4
//	Loop Iteration 5
//	Countdown 10
//	Countdown 9
//	Countdown 8
//	Countdown 7
//	Countdown 6
//	Countdown 5
//	Countdown 4
//	Countdown 3
//	Countdown 2
//	Countdown 1
//	Lift Off!
