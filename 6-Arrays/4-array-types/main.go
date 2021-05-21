package main

import (
	"fmt"
)

// array types (when passed to a function)

// an "array" contains a sequence of "values"
// a "slice" contains a "reference" to a sequence of values

// when passing a type to a function:
// * types that are 'value-based' (ARRAY) receive a COPY of the original
// * types that are 'references' (SLICE) point to THE original
// * in order to operate on an original 'value', POINTERS are needed when passing value-based (ARRAY) types

// 'variadic' function accepts any number of arguments and all these arguments are stored in a parameter of the slice type
// the arguments are stored in a parameter of the `slice` type
// called with any number of arguments
// func f(elem ...Type)
// `...` 'pack operator' instructs Go to store all arguments of type `Type` in `elem` parameter
func list(groceries ...string) {

	for i := 0; i < len(groceries); i++ {
		fmt.Printf("\n%v. %v", i, groceries[i])
	}
	fmt.Println()
}

func main() {

	// make array
	array := [4]string{"coffee", "bananas", "pasta", "eggs"}

	// copy array elements into a `slice` named 'slicedArray'
	// variadic functions require array element values be converted to a slice 
	slicedArray := array[:]

	// call 'variadic' function `list` and pass `slicedArray` array element values
	// `...` unpack operator -unpack a slice
	list(slicedArray...)

	// add 2 values to `slice` array using 'append' function
	slicedArray = append(slicedArray, "hot sauce", "steak sauce")

	list(slicedArray...)
}

//	% go run main.go
//	
//	0. coffee
//	1. bananas
//	2. pasta
//	3. eggs
//	
//	0. coffee
//	1. bananas
//	2. pasta
//	3. eggs
//	4. hot sauce
//	5. steak sauce
