package main

import (
	"fmt"
)

// a slice can be independently created

// use same syntax as array except minus the number of elements specified
// ----------------------------------
// sliceName := []dataType{value, value, value,} // intialized 'slice literal'

// `slice` type contains 3 components:
//  1) `Ptr` - (pointer) a pointer to the 1st element in the array
//  2) `Len` - (length) the total number of elements in the slice
//  3) `Cap` - (capacity) the number of elements between the element referenced by the pointer and the last array element

// use built-in function `make()`
// ----------------------------------
// hondas := make([]string, 3)
// hondas[0] = "Civic"
// hondas[1] = "Accord"
// hondas[2] = "CR-V"


func sliceTypeComponents(foreignAutosSlice []string) {
	fmt.Printf("\n'foreignAutosSlice': %v ", foreignAutosSlice)
	fmt.Printf("\n'Ptr': %p ", foreignAutosSlice)
	fmt.Printf("\n'Len': %v ", len(foreignAutosSlice))
	fmt.Printf("\n'Cap': %v \n", cap(foreignAutosSlice))
}

func main() {

	// create an initialized slice literal
	foreignAutosSlice := []string{"Sentra", "Accent", "Sonata", "Jetta", "Corolla"}

	// create empty slice `hondasSlice`
	hondasSlice := make([]string, 3)
	hondasSlice[0] = "Civic"
	hondasSlice[1] = "Accord"
	hondasSlice[2] = "CR-V"
	
	sliceTypeComponents(foreignAutosSlice)

	// add element to `foreignAutosSlice` using `append()` function
	// `append()` returns a new slice containing additional elements
	// *** when length of the appended slice exceeds the 'Capacity' of base array, Go increases 'Capacity' of the array ***
	// *** Go automatically increses 'Capacity' and it may double the 'Capacity' number even if greater than actual number of elements ***
	foreignAutosSlice = append(foreignAutosSlice, hondasSlice[0])
	sliceTypeComponents(foreignAutosSlice)

	// built-in `copy()` function copies one slice into another
	// `copy()` function takes 2 args (slice to copy into, slice to copy from)
	// technique to remove an element from a slice and preserve order of remaining elements
	// below, `copy()` function is removing 2st element of slice `foreignAutosSlice[1:]` by copying in slice `foreignAutosSlice[2:]`
	copy(foreignAutosSlice[1:], foreignAutosSlice[2:])
	// order is maintained (appended Civic will be last) but need to remove last value of duplicated 'Civic'
	foreignAutosSlice = foreignAutosSlice[:len(foreignAutosSlice)-1]
	sliceTypeComponents(foreignAutosSlice)
}

//	% go run main.go
//	
//	'foreignAutosSlice': [Sentra Accent Sonata Jetta Corolla] 
//	'Ptr': 0xc000068050 
//	'Len': 5 
//	'Cap': 5 
//	
//	'foreignAutosSlice': [Sentra Accent Sonata Jetta Corolla Civic] 
//	'Ptr': 0xc000106000 
//	'Len': 6 
//	'Cap': 10 
//	
//	'foreignAutosSlice': [Sentra Sonata Jetta Corolla Civic] 
//	'Ptr': 0xc000106000 
//	'Len': 5 
//	'Cap': 10
