package main

import "fmt"

func main() {
	i := 1

	// Declaring pointer
	var p *int = nil

	p = &i // Getting address of i

	fmt.Println(p)  // memory address
	fmt.Println(*p) // pointer value

	*p++            // unreferencing, getting the value
	fmt.Println(p)  // memory address
	fmt.Println(*p) // pointer value

	i++
	fmt.Println(p)  // memory address
	fmt.Println(*p) // pointer value
	fmt.Println(p, *p, i, &i)

	/*
		There is no arithmetical operation of pointers on Golang.
		like p++
	*/
}
