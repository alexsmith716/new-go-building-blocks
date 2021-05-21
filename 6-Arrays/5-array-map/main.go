package main

import (
	"fmt"
)

// array map data structure
// review/as expected

// data can be stored with associated keys
// * map 'keys' are unique and same data type
// * map 'values' are unique and same data type

// var mapName = map[ keyDataType ]valueDataType
// mapName[ "keyName" ] = value

func main() {

	// using Go's `make` function, create a map for string keys and values
	// capitalCity := make(map[string]string)
	// capitalCity["England"] = "London"
	// capitalCity["France"] = "Paris"
	// capitalCity["Germany"] = "Berlin"

	// map literal
	capitalCity := map[string]string {
		"England": "London",
		"France": "Paris",
		"Germany": "Berlin",
	}

	fmt.Printf("capitalCity: %v \n", capitalCity)

	capitalCity["Spain"] = "Madrid"

	fmt.Printf("Spain capitalCity: %v \n", capitalCity["Spain"])

	delete(capitalCity, "Spain")

	for i, c := range capitalCity {
		fmt.Printf("capitalCity %v : %v \n", i, c)
	}
}

//	% go run main.go
//	capitalCity: map[England:London France:Paris Germany:Berlin] 
//	Spain capitalCity: Madrid 
//	capitalCity England : London 
//	capitalCity France : Paris 
//	capitalCity Germany : Berlin
