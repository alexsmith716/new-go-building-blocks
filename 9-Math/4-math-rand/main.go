package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 'math' provides 'rand' package to generate (relatively) random numbers
// functions:
// `rand.Intn(44)` returns a random number from 0 up to 44 (argument)
// `rand.Perm(26)` returns a sequence of integers, from 0 up to 26 (argument)
// `rand.Seed()` 
// the randomizing pattern is based a default 'seed' which always returns the same 'random' order
// the 'seed' can be specified which generates a different order (using system time function `time.Now()`)
// using function `time.Now()`, the nanoseconds are extracted (`UnixNano()`) and used as the 'seed'


func main() {

	// specify system time 'nanoseconds' as a randomizing pattern 'seed'
	rand.Seed(time.Now().UnixNano())

	// return a random number between 0 and 20
	randIntn := rand.Intn(20)

	fmt.Println("rand.Intn(20): ", randIntn)

	// return a sequence of random numbers between 0 and 20
	randPerm := rand.Perm(20)

	fmt.Println("rand.Perm(20): ", randPerm)
}

//	% go run main.go
//	rand.Intn(20):  8
//	rand.Perm(20):  [9 8 3 13 5 10 1 17 15 12 0 2 11 7 14 19 4 18 16 6]
//	% go run main.go
//	rand.Intn(20):  3
//	rand.Perm(20):  [15 11 16 13 18 5 14 10 7 8 2 9 4 3 17 1 0 19 12 6]
//	% go run main.go
//	rand.Intn(20):  13
//	rand.Perm(20):  [7 19 9 14 6 15 10 3 2 4 1 16 12 18 17 0 8 5 11 13]
//	% go run main.go
//	rand.Intn(20):  8
//	rand.Perm(20):  [7 11 16 13 17 9 2 19 5 18 15 14 4 6 3 0 12 10 1 8]

