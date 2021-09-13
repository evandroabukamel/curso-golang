package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var ray = 3.2 // type infered is float64 by compiler

	// Reduced form to create a variable
	area := PI * m.Pow(ray, 2)

	fmt.Println("The area of circle is", area)

	// Another way to declare const and var
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Booleans
	var e, f bool = true, false
	fmt.Println(e, f)

	// Multiple types at once
	g, h, i := 2, false, "oops"
	fmt.Println(g, h, i)
}
