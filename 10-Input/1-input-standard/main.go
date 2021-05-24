package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 'fmt' package provides functions () to read from standard imput:

// Scan(), ScanIn(), Scanf()

// * each function returns 2 values, the number of items scanned and read errors
// * each function stores successive space-separated values into successive interface variables


func main() {

	// specify system time 'nanoseconds' as a randomizing pattern 'seed'
	rand.Seed(time.Now().UnixNano())

	// return a random number between 0 and 20 (minding zero based "+1")
	var randIntn int = rand.Intn(20) + 1
	var usersGuess int = 0
	var flag bool = true

	fmt.Print("guess a random number between 1-20: ")

	for flag {

		// compare `usersGuess` input to 'randIntn'
		// for function 'Scan()', specify the address "&" of interface variable `usersGuess`
		// numberOfItemsScanned, err := fmt.Scan(&usersGuess)
		_, err := fmt.Scan(&usersGuess)

		if err != nil {
			fmt.Println(err)
		} else if usersGuess > randIntn {
			fmt.Print("number is lower, try again: ")
		} else if usersGuess < randIntn {
			fmt.Print("number is higher, try again: ")
		} else if usersGuess == randIntn {
			fmt.Println("guessed correct number:",randIntn)
			flag = false
		}
	}
}

//	% go run main.go
//	guess a random number between 1-20: 9
//	number is higher, try again: 10
//	number is higher, try again: 11
//	number is higher, try again: 12
//	guessed correct number: 12
