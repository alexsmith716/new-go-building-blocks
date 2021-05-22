package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// create `Time` data types using `Date()` function
// requires 8 arguments to specify:
// (year, month, day number, hour, minute, second, nanosecond, location)

// USE THIS TIME STRING FORMAT: "Mon Jan 2 15:04:05 2006 MST"


func main() {

	timeDate := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.Local)
	fmt.Printf("time.Date(): %v \n", timeDate)

	timeDate = timeDate.AddDate(2, 6, 3)
	fmt.Printf("timeDate.AddDate(): %v \n", timeDate)

	// create `Time` data type from a layout string using `Parse()` function
	layout := "2006-Jan-02 15:04PM"
	str := "2100-Jan-01 12:01AM"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("time.Parse(): %v \n", t)
	}
}

//	% go run main.go
//	time.Date(): 2025-01-01 12:00:00 -0500 EST 
//	timeDate.AddDate(): 2027-07-04 12:00:00 -0400 EDT 
//	time.Parse(): 2100-01-01 00:01:00 +0000 UTC
