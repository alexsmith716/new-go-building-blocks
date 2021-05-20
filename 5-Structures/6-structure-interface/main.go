package main

import (
	"fmt"
)

// Go is a type-safe, statically typed, compiled programming language
// Go's type system specifies a set of rules that assign a type property to variables, functions and identifiers
// types (type names / type declarations) are designed to prevent occurrences of unchecked runtime type errors

// represent objects in programs -OOP
// Encapsulation, Inheritance and Polymorphism is OOP

// Polymorphism:
// it means many forms
// its the ability to assign a different meaning or purpose to an object according to the object's context
// polymorphism is supported through an "interface" type
// an interface defines and describes the exact methods that some other type must have
// "interfaces reduce duplication or boilerplate code"
// "interfaces make it easier to use mocks instead of real objects in unit tests"
// "interfaces enforce decoupling between parts of the codebase"
// "interfaces do not enforce a type, a type can choose to implement methods of an interface"

// interface body contains list of method signitures
//
//	type interfaceName interface {
//		methodName returnType
//		methodNme returnType
//	}

// any struct that has methods that match methods in interface will implement that interface

// interface Vehicle | interface Aircraft

type aircraft interface {
	liftType() string
	propulsionType() string
}

// ------------------------------------

type cessna172 struct{}

func (cessna172) liftType() string {
	return "dynamic"
}

func (cessna172) propulsionType() string {
	return "propeller"
}

// ------------------------------------

type hotAirBalloon struct{}

func (hotAirBalloon) liftType() string {
	return "static"
}

func (hotAirBalloon) propulsionType() string {
	return "wind"
}

// ------------------------------------

func main() {

	// interfaces in Go are "implicitly implemented"
	// type `cessna172` struct type implements interface `aircraft`
	// so, a variable of type `cessna172` can be represented as the type of interface `aircraft`
	var a1 aircraft
	a1 = cessna172{}
	fmt.Printf("%v \n", a1.liftType())
	fmt.Printf("%v \n", a1.propulsionType())

	// type `hotAirBalloon` struct type implements interface `aircraft`
	var a2 aircraft
	a2 = hotAirBalloon{}
	fmt.Printf("%v \n", a2.liftType())
	fmt.Printf("%v \n", a2.propulsionType())
}

//	% go run mainC.go
//	dynamic 
//	propeller 
//	static 
//	wind
