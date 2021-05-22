package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// `time` package: extract fields from `Time` data type that describe a point in time

// individual 'time' value fields of `Time` data type

func main() {

	// `time.Now()` function instance describing the current date & time

	// methods returned by `Time` data type (methods return individual 'time' value fields)
	// "Hour()" day hour number (0-23) 
	// "Minute()" hour minute number (0-59)
	// "Second()" minute second number (0-59)
	// "Nanosecond()" nanosecond of the second (0-99999)
	// "Clock()" day hour, minute, and second
	// "Unix()" seconds elapsed since epoch
	// "UnixNano()" nanoseconds since epoch

	timeNow := time.Now()

	fmt.Printf("time.Now(): %v \n", timeNow)
	fmt.Printf("time.Now() Type: %T \n", timeNow)

	hour := timeNow.Hour()

	switch {
	case hour < 12:
		fmt.Println("hour is morning")
	case hour < 18:
		fmt.Println("hour is afternoon")
	default:
		fmt.Println("hour is night")

	}

	h, mn, s := timeNow.Clock()
	fmt.Printf("timeNow.Clock(): %v:%v:%v \n", h, mn, s)

	ns := timeNow.UnixNano()
	ms := ns / 1000000
	fmt.Println("nanoseconds:", ns)
	fmt.Println("milliseconds:", ms)
}

//	% go run main.go
//	time.Now(): 2021-05-21 15:03:59.628721 -0400 EDT m=+0.000064246 
//	time.Now() Type: time.Time 
//	hour is afternoon
//	timeNow.Clock(): 15:3:59 
//	nanoseconds: 1621623839628721000
//	milliseconds: 1621623839628
