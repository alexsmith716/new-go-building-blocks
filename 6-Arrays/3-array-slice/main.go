package main

import (
	"fmt"
)

// slicing arrays
// review/as expected

// `slice` type represents an array whose length can be changed
// `slice` is a reference to an actual array

// `slice` type contains 3 components:
//	1) `Ptr` - (pointer) a pointer to the 1st element in the array
//	2) `Len` - (length) the total number of elements in the slice
//	3) `Cap` - (capacity) the number of elements between the element referenced by the pointer and the last array element

//  a `slice` is created by assigning an array to a `var` that specifies element index numbers of a desired lower & upper boundry
//  weekdays := days[0:5]

func main() {

	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	weekdays := days[0:5]

	weekend := days[5:]

	fmt.Printf("slice weekdays: %v \n", weekdays)

	fmt.Printf("slice weekend: %v \n", weekend)
	fmt.Printf("type weekend : %T \n", weekend)

	fmt.Printf("Len weekend: %v \n", len(weekend))
	fmt.Printf("Cap weekend: %v \n", cap(weekend))

	fmt.Printf("Ptr weekend: %p \n", weekend)
	fmt.Printf("Address days[0]: %p \n", &days[5])
}

//	% go run main.go
//	slice weekdays: [Mon Tue Wed Thu Fri] 
//	slice weekend: [Sat Sun] 
//	type weekend : []string 
//	Len weekend: 2 
//	Cap weekend: 2 
//	Ptr weekend: 0xc0000b6050 
//	Address days[0]: 0xc0000b6050
