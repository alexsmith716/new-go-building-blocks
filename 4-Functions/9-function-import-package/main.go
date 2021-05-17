package main

import (
	"fmt"
	"new-go-building-blocks/utils"
)

// import separate/custom package (mind the upper/lower cases)
// review/as expected


func main() {

	for i := 2; i >= -2; i-- {

		res, err := utils.IsAPositiveInteger(i)

		if err != nil {
			fmt.Println("Failed:", err)
		} else {
			fmt.Println(res, "Passed!")
		}
	}
}

//	% go run main.go
//	2 Passed!
//	1 Passed!
//	Failed: 0 not a positive integer
//	Failed: -1 not a positive integer
//	Failed: -2 not a positive integer
