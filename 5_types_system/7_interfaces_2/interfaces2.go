package main

import "fmt"

type Sport interface {
	turnOnTurbo()
}

type Ferrari struct {
	model        string
	isTurboOn    bool
	currentSpeed int
}

func (f *Ferrari) turnoOnTurbo() {
	f.isTurboOn = true
}

func main() {
	car1 := Ferrari{"F40", false, 0}
	car1.turnoOnTurbo()
	fmt.Println(car1)

	// Not working anymore
	/*var car2 Sport = &Ferrari{"F50", false, 0}
	car2.turnOnTurbo()
	fmt.Println(car2)*/
}
