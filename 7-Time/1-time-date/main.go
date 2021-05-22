package main

import (
	"fmt"
	"time"
)

// Go standard library `time` package
// `time` package: extract fields from `Time` data type that describe a point in time

func main() {

	// `time.Now()` function instance describing the current date & time

	// methods returned by `Time` data type (methods return individual 'date' value fields)
	// "Now()" the `time.Time` data type
	// "Year()" year with 4 digits
	// "Month()" month name of the year
	// "Day()" day number of the month
	// "Weekday()" day name of the week
	// "Date()" year (2021)
	// "ISOWeek()" week number of the year (1-53)
	// "YearDay()" day number of the year (1-365)

	timeNow := time.Now()

	fmt.Printf("time.Now(): %v \n", timeNow)
	fmt.Printf("time.Now() Type: %T \n", timeNow)

	fmt.Printf("timeNow.Weekday() is: %v \n", timeNow.Weekday())

	y, m, d := timeNow.Date()
	fmt.Printf("timeNow.Date(): %v %v, %v \n", m, d, y)

	year, week := timeNow.ISOWeek()
	fmt.Printf("timeNow.ISOWeek(): %v in %v\n", week, year)

	day := timeNow.YearDay()
	fmt.Printf("timeNow.YearDay(): %v \n", day)
}

//	% go run main.go
//	time.Now(): 2021-05-21 14:23:18.098503 -0400 EDT m=+0.000065643 
//	time.Now() Type: time.Time 
//	timeNow.Weekday() is: Friday 
//	timeNow.Date(): May 21, 2021 
//	timeNow.ISOWeek(): 20 in 2021
//	timeNow.YearDay(): 141
