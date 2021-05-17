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

//	% go run main.go
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
//	Factorial 1 = 1
//	Factorial 2 = 2
//	Factorial 3 = 6
//	Factorial 4 = 24
//	Factorial 5 = 120
//	Factorial 6 = 720
//	Factorial 7 = 5040
