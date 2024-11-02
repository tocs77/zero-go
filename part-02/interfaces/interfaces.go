//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	LiftUp() string
}

type Car string

func (c Car) LiftUp() string {
	return "Standard Lift"
}

type Truck string

func (t Truck) LiftUp() string {
	return "Large Lift"
}

type Motorcycle string

func (m Motorcycle) LiftUp() string {
	return "Small Lift"
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("Super puper Motorcycle: %s", string(m))
}

func liftVeicles(l []Lifter) {
	for _, lift := range l {
		fmt.Println(lift.LiftUp())
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")
	liftVeicles([]Lifter{car, truck, motorcycle})

	fmt.Println(motorcycle)
}
