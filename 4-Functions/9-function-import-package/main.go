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
