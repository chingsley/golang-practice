package interfaces

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

func Execute() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says: ", a.Says())
	fmt.Println("This animal has: ", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Uhh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
