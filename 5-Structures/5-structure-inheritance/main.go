package main

import (
	"fmt"
	"strconv" 
)

// https://golang.org/pkg/strconv/

// represent objects in programs -OOP
// Encapsulation, Inheritance and Polymorphism is OOP

// Inheritance:
// objectRaccoon (derived from objectMammal) will inherit the features of objectMammal
// "composition" is the code reuse of an object derived from another object
// "composition" is the embedding of one struct inside a 2nd struct which provides the 2nd the elements of the 1st
// the fields and attached methods are "made accessible/inherited" to the 2nd/inherited struct

type cargoShip struct {
	name string
	inService bool
	containerCapacity int
}

func (c cargoShip) describeCargoShip() string {
	return c.name + " | " + "Service: " + strconv.FormatBool(c.inService) + " | " + "Capacity: " + strconv.Itoa(c.containerCapacity)
}

type itinerary struct {
	detination string
	cargoType string
	totalContainers int
	cargoShip
}

func (i itinerary) listItinerary() {
	fmt.Println("Detination:", i.detination)
	fmt.Println("CargoType:", i.cargoType)
	fmt.Println("TotalContainers:", i.totalContainers)
	fmt.Printf("Vessel: %v \n", i.describeCargoShip())
	fmt.Println("-------------------------------------------")
}

func main() {

	cargoShip1 := cargoShip{
		"Grand Cargo",
		true,
		3500,
	}

	itinerary1 := itinerary{
		"Seattle",
		"textile",
		3400,
		cargoShip1,
	}
	itinerary1.listItinerary()

	itinerary2 := itinerary{
		"New York",
		"perishable",
		1250,
		cargoShip1,
	}
	itinerary2.listItinerary()
}

//	% go run mainA.go
//	Detination: Seattle
//	CargoType: textile
//	TotalContainers: 3400
//	Vessel: Grand Cargo | Service: true | Capacity: 3500 
//	-------------------------------------------
//	Detination: New York
//	CargoType: perishable
//	TotalContainers: 1250
//	Vessel: Grand Cargo | Service: true | Capacity: 3500 
//	-------------------------------------------
