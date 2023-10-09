package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says ", a.Says(), "and has ",
		a.NumberOfLegs(), " legs.")
}

func (d *Dog) Says() string {
	return "Woof!"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Eee Eee, Ooh Ooh!"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
func main() {
	dog := Dog{
		Name:  "Richard",
		Breed: "German Shepard",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}
