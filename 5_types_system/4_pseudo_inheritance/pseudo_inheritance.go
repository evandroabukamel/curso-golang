package main

import "fmt"

type Car struct {
	name         string
	currentSpeed int
}

type Ferrari struct {
	Car     // hidden fields
	turboOn bool
}

func main() {
	f := Ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Printf("Is Ferrari %s turbo on? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}
