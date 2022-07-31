package main

import "fmt"

type Sporting interface {
	turboOn()
}

type LX interface {
	park()
}

type SportingLX interface {
	Sporting
	LX
}

type BMW7 struct{}

func (c BMW7) turboOn() {
	fmt.Println("Turbo on...")
}

func (c BMW7) park() {
	fmt.Println("Start parking...")
}

func main() {
	var b SportingLX = BMW7{}
	b.turboOn()
	b.park()
}
