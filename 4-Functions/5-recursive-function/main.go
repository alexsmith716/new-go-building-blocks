package main

import (
	"fmt"
)

// recursive function
// review/as expected

// functions calling themselves (and modifying a tested expression & eventually exit)


func countDn(num int) {
	if num < 1 {
		fmt.Println("Lift Off!")
	} else {
		fmt.Println("Countdown", num)
		num--
		countDn(num)
	}
}

func facto(num int) int {

	if num == 0 {
		return 1
	}
	return num * facto(num-1)
}


func main() {
	countDn(10)

	for i := 1; i <= 7; i++ {
		fmt.Println("Factorial", i, "=", facto(i))
	}
}
