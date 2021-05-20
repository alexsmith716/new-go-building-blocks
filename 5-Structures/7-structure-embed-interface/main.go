package main

import (
	"fmt"
)

// type interface (embedded)

// Go is a type-safe, statically typed, compiled programming language
// Go's type system specifies a set of rules that assign a type property to variables, functions and identifiers
// types (type names / type declarations) are designed to prevent occurrences of unchecked runtime type errors

// represent objects in programs -OOP
// Encapsulation, Inheritance and Polymorphism is OOP

// structs can be embedded and interface can be embedded
// Go automatically identifies which struct implements each interface
// interface embedding common in Go standard libary
// the `io` package contains a `ReadWriter` interface that embeds `Reader` and `Writer` interfaces
// those embedded `Reader` and `Writer` interfaces contain `Read` and `Write` method signitures

// an interface cannot implement or extend other interfaces
// create a new interface by merging two or more interfaces

// land, rail, water, amphibious, air
// internal combustion, battery, propeller, jet, wind, nuclear
// internal combustion, battery, propeller, jet, wind, nuclear

type transportation interface {
	transportationType() string
}

type propulsion interface {
	propulsionType() string
	propulsionMethod() string
}

// vehicle: object that transports people or other objects
// new interface `vehicle` with merged interface `transportation` & `propulsion`
type vehicle interface {
	transportation
	propulsion
}

// ------------------------------------

type subwayCar struct{}

func (subwayCar) transportationType() string {
	return "rail"
}

func (subwayCar) propulsionType() string {
	return "electric"
}

func (subwayCar) propulsionMethod() string {
	return "self propelled"
}

// ------------------------------------

type cargoShip struct{}

func (cargoShip) transportationType() string {
	return "water"
}

func (cargoShip) propulsionType() string {
	return "internal combustion"
}

func (cargoShip) propulsionMethod() string {
	return "propeller thrust"
}

// ------------------------------------

func main() {
	// since `subwayCar` implements methods `transportationType` & `propulsionType` & `propulsionMethod`, it implements interfaces `transportation` and `propulsion`
	// since interface `vehicle` is an embedded interface of `transportation` and `propulsion`, `subwayCar` also implements `vehicle`
	s := subwayCar{}
	var v1 vehicle = s
	var t1 transportation = s
	var p1 propulsion = s
	fmt.Printf("dynamic type of interface v1 of static type `vehicle` is '%T' \n", v1)
	fmt.Printf("dynamic type of interface t1 of static type `transportation` is '%T' \n", t1)
	fmt.Printf("dynamic type of interface p1 of static type `propulsion` is '%T' \n", p1)

	c := cargoShip{}
	var v2 vehicle = c
	var t2 transportation = c
	var p2 propulsion = c
	fmt.Printf("dynamic type of interface v2 of static type `vehicle` is '%T' \n", v2)
	fmt.Printf("dynamic type of interface t2 of static type `transportation` is '%T' \n", t2)
	fmt.Printf("dynamic type of interface p2 of static type `propulsion` is '%T' \n", p2)
}

//	% go run main.go
//	dynamic type of interface v1 of static type `vehicle` is 'main.subwayCar' 
//	dynamic type of interface t1 of static type `transportation` is 'main.subwayCar' 
//	dynamic type of interface p1 of static type `propulsion` is 'main.subwayCar' 
//	dynamic type of interface v2 of static type `vehicle` is 'main.cargoShip' 
//	dynamic type of interface t2 of static type `transportation` is 'main.cargoShip' 
//	dynamic type of interface p2 of static type `propulsion` is 'main.cargoShip'
