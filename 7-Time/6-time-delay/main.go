package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// `Sleep()` function allows delaying progress of a program by specified amount

// USE THIS TIME STRING FORMAT: "Mon Jan 2 15:04:05 2006 MST"

// `time` package has 3 methods that compare an instance of `Time` data type
// * "Before(time)" returns a boolean `true` value if the instance time is earlier than argument time
// * "After(time)"  returns a boolean `true` value if the instance time is later than argument time
// * "Equal(time)"  returns a boolean `true` value if the instance time and argument time are same

// delaying execution is useful in pausing individual 'goroutines'

func main() {

	startTime := time.Now()
	fmt.Println("startTime:", startTime.Format("15:04:05"))
	
	// delay program progress for 6 seconds
	time.Sleep(6 * time.Second)

	finishTime := time.Now()
	fmt.Println("finishTime:", finishTime.Format("15:04:05"))

	fmt.Println("startTime.Before(finishTime):", startTime.Before(finishTime))
	fmt.Println("finishTime.Before(startTime):", finishTime.Before(startTime))

	// `Sub()` method returns difference between 2 instances of `Time` data type
	elapsed := finishTime.Sub(startTime)
	// `Round()` method rounds a duration to the nearest specified amount (time.Second)
	fmt.Println("diff.Round(time.Second):", elapsed.Round(time.Second))
}

//	% go run main.go
//	startTime: 21:05:45
//	finishTime: 21:05:51
//	startTime.Before(finishTime): true
//	finishTime.Before(startTime): false
//	diff.Round(time.Second): 6s
