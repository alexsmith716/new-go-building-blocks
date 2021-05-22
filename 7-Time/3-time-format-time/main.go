package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// `time` package: extract fields from `Time` data type that describe a point in time

// individual fields of date and time values can be formatted using `Format()` method
// `Format()` method takes a 'layout' string
// 'layout' string specifies how the to format a date/time value

// USE THIS TIME STRING FORMAT: "Mon Jan 2 15:04:05 2006 MST"

func main() {

	timeNow := time.Now()

	fmt.Println("time.Now():", timeNow)
	fmt.Println("timeNow.Format(time.UnixDate):", timeNow.Format(time.UnixDate))
	fmt.Println("timeNow.Format(time.ANSIC):", timeNow.Format(time.ANSIC))
	fmt.Println("timeNow.Format(time.RFC3339):", timeNow.Format(time.RFC3339))
	fmt.Println("timeNow.Format('January 2, 2006 [Monday]'):", timeNow.Format("January 2, 2006 [Monday]"))

	fmt.Println("timeNow.Format('January 2, 2006'):", timeNow.Format("January 2, 2006"))
	fmt.Println("timeNow.Format('2 January, 2006'):", timeNow.Format("2 January, 2006"))

	// `time.Kitchen` constant has layout format of "12:11PM"
	fmt.Println("timeNow.Format(time.Kitchen):", timeNow.Format(time.Kitchen))
	fmt.Println("timeNow.Format('15:04'):", timeNow.Format("15:04"))
}

//	% go run main.go
//	time.Now(): 2021-05-21 16:01:32.732662 -0400 EDT m=+0.000071516
//	timeNow.Format(time.UnixDate): Fri May 21 16:01:32 EDT 2021
//	timeNow.Format(time.ANSIC): Fri May 21 16:01:32 2021
//	timeNow.Format(time.RFC3339): 2021-05-21T16:01:32-04:00
//	timeNow.Format('January 2, 2006 [Monday]'): May 21, 2021 [Friday]
//	timeNow.Format('January 2, 2006'): May 21, 2021
//	timeNow.Format('2 January, 2006'): 21 May, 2021
//	timeNow.Format(time.Kitchen): 4:01PM
//	timeNow.Format('15:04'): 16:01
