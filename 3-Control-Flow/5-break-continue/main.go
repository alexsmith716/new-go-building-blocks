package main

import (
	"fmt"
)

// while-loop
// review/as expected

func main() {

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			if i == 3 && j == 2 {
				fmt.Println("Continues When i=", i, "and j=", j)
				continue
			}

			if i == 2 && j == 2 {
				fmt.Println("Breaks When i=", i, "and j=", j)
				break
			}

			fmt.Println("Running i=", i, "j=", j)

		}
	}
}

//	% go run main.go
//	Running i= 1 j= 1
//	Running i= 1 j= 2
//	Running i= 1 j= 3
//	Running i= 2 j= 1
//	Breaks When i= 2 and j= 2
//	Running i= 3 j= 1
//	Continues When i= 3 and j= 2
//	Running i= 3 j= 3
