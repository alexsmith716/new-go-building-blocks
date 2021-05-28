package main

import (
	"fmt"
)

// slicing arrays

// >>>>>>>>>>> an ARRAY contains a sequence of values
// >>>>>>>>>>> a SLICE contains a REFERENCE to a sequence of values (array elements)
// (slice has a 'pointer' to an array element 'address')

// `slice` type contains 3 components:
//	1) `Ptr` - (pointer) a pointer to the 1st element in the array
//	2) `Len` - (length) the total number of elements in the slice
//	3) `Cap` - (capacity) the number of elements between the element referenced by the pointer and the last array element

//  a `slice` is created by assigning an array to a `var` that specifies element index numbers of a desired lower & upper boundry
//  sliceWeekDays := arrayDaysOfTheWeek[0:5]

func main() {

	// create array of 7 elements
	arrayDaysOfTheWeek := [7]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}

	fmt.Printf("array arrayDaysOfTheWeek: %v \n", arrayDaysOfTheWeek)
	fmt.Printf("type arrayDaysOfTheWeek : %T \n", arrayDaysOfTheWeek)

	// create a slice of weekdays from the array (that represents the 5 'sliceWeekDays')
	// 'Ptr' starts at 0 and the 'Len'/number is 5 elements
	sliceWeekDays := arrayDaysOfTheWeek[0:5]

	// create a slice of weekdays from the array (that represents the 2 'sliceWeekend' days)
	// 'Ptr' starts at 5 and the 'Len'/number is omitted so its the 
	sliceWeekend := arrayDaysOfTheWeek[5:]

	sliceMonday := arrayDaysOfTheWeek[0:1]

	fmt.Printf("sliceMonday: %v \n", sliceMonday)
	fmt.Printf("'Ptr' sliceMonday: %p \n", sliceMonday)
	fmt.Printf("Address &arrayDaysOfTheWeek[0]: %p \n", &arrayDaysOfTheWeek[0])

	fmt.Printf("arrayDaysOfTheWeek[0]: %v \n", arrayDaysOfTheWeek[0])
	fmt.Printf("&arrayDaysOfTheWeek[0]: %v \n", &arrayDaysOfTheWeek[0])

	fmt.Printf("arrayDaysOfTheWeek[5]: %v \n", arrayDaysOfTheWeek[5])
	fmt.Printf("&arrayDaysOfTheWeek[5]: %v \n", &arrayDaysOfTheWeek[5])

	fmt.Printf("arrayDaysOfTheWeek[6]: %v \n", arrayDaysOfTheWeek[6])
	fmt.Printf("&arrayDaysOfTheWeek[6]: %v \n", &arrayDaysOfTheWeek[6])

	fmt.Printf("slice sliceWeekDays: %v \n", sliceWeekDays)

	fmt.Printf("slice sliceWeekend: %v \n", sliceWeekend)
	fmt.Printf("type sliceWeekend : %T \n", sliceWeekend)

	fmt.Printf("'Len' sliceWeekend: %v \n", len(sliceWeekend))
	fmt.Printf("'Cap' sliceWeekend: %v \n", cap(sliceWeekend))

	// confirm that slice 'sliceWeekend' is a reference pointing to array element 'arrayDaysOfTheWeek[5]'
	// compare the pointer value of slice 'sliceWeekend' with the memory address '&' of the array element that it points to
	fmt.Printf("'Ptr' sliceWeekend: %p \n", sliceWeekend)
	fmt.Printf("Address arrayDaysOfTheWeek[5]: %p \n", &arrayDaysOfTheWeek[5])
}

//	% go run main.go
//	slice sliceWeekDays: [Mon Tue Wed Thu Fri] 
//	slice sliceWeekend: [Sat Sun] 
//	type sliceWeekend : []string 
//	Len sliceWeekend: 2 
//	Cap sliceWeekend: 2 
//	Ptr sliceWeekend: 0xc0000b6050 
//	Address arrayDaysOfTheWeek[0]: 0xc0000b6050
