package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name           string
	Colour         string
	NumberOfTeeths int
}

// ==================== MAIN FUNCTION ====================
func main() {

	dog := Dog{
		Name:  "Deo",
		Breed: "Labrador",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:           "Jock",
		Colour:         "Black",
		NumberOfTeeths: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {

	log.Println("This animals says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
