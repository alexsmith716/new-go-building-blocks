package main

import (
	"fmt"
)

// a struct is a type
// a struct 'field' is declared with a 'name' and 'type'
// embed one struct within another by declaring a 'field' to be a 'struct' type

//	type structName struct {
//		fieldName embeddedStructName
//	}

type coordinates struct {
	x, y int
}

type circle struct {
	radius int
	// center coordinates // works
	coordinates
}


// attach method `methodDiameter()` to struct `circle`
func (c circle) methodDiameter() int {
	return c.radius * 2
}


func main() {

	// instance of `circle`
	var ring circle

	// fields of embedded struct `coordinates` can be accessed
	ring.radius = 35
	ring.x = ring.radius
	ring.y = ring.radius

	// or, fields of embedded struct `coordinates` can be accessed by chaining
	// instanceName.embeddedStructName.fieldName
	// ring.coordinates.x = ring.radius
	// ring.coordinates.y = ring.radius
	// ring.center.x = ring.radius
	// ring.center.y = ring.radius

	fmt.Printf("Diameter:%v\n", ring.methodDiameter())
	//	fmt.Printf("Point X:%v Y:%v \n", ring.center.x, ring.center.y) // works
	fmt.Printf("Point X:%v Y:%v \n", ring.x, ring.y)
}

//	% go run main.go
//	Diameter:70
//	Point X:35 Y:35
