package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// all `Time` data types are associated with a time zone (UTC time/local time)
// specify an alternative location/time zone using function `time.LoadLocation()`

func main() {

	tZone := "All"

	// `time` data type in specified location "America/New_York"
	// `time.LoadLocation()` function returns a 'Location' type and an error
	// on successful function return error value is `nil`, otherwise error message
	// "_" blank identifier used to ignore returned error value
	location, _ := time.LoadLocation("America/New_York") // <<< IANA time zone

	// create a `Time` data type of current time in different time zone
	// append a call to `In()` function after `Now()` function
	// the location is specified as arg to `In()`
	timeNow := time.Now().In(location)

	// get the time zone associated with `Time` data type with `Zone()` method
	// `Zone()` method returns a 3-letter TZ abbreviation and the seconds offset of that TZ from UTC
	// convert the seconds offset to a positive minute number
	tzAbbreviation, tzOffset := timeNow.Zone()
	fmt.Println("tzOffset: ", tzOffset)
	if tzOffset < 1 {
		tzOffset = tzOffset * -1
	}
	tzOffset /= 60

	// evaluate the abbreviation and offset values
	switch {
	case tzAbbreviation == "EST" && tzOffset == 300:
		tZone = "East Coast"
	case tzAbbreviation == "EDT" && tzOffset == 240:
		tZone = "East Coast"
	}

	fmt.Printf("tzAbbreviation: %v tzOffset: %v \n", tzAbbreviation, tzOffset)
	fmt.Println("tZone: ", tZone)
}

//	% go run main.go
//	tzOffset:  -14400
//	tzAbbreviation: EDT tzOffset: 240 
//	tZone:  East Coast
