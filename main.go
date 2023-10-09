package main

import "fmt"

// Declare a Animal interface
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Create 2 different structs
type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Create a function that will print some output when we run the program.
func PrintInfo(a Animal) {
	fmt.Println("This animal says ", a.Says(), "and has ",
		a.NumberOfLegs(), " legs.")
}

// Use the * to denote pointers to the d varriable
func (d *Dog) Says() string {
	return "Woof!"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Use the * to denote pointers to the d varriable
func (d *Gorilla) Says() string {
	return "Eee Eee, Ooh Ooh!"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
func main() {

	// Create a new dog variable using or dog Struct
	// Call the PrintInfo Function
	dog := Dog{
		Name:  "Richard",
		Breed: "German Shepard",
	}
	PrintInfo(&dog)
	// Create a new dog variable using or dog Struct
	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}
