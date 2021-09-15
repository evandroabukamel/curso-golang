package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("line.")

	x := 3.141516
	// fmt.Println("The value of x is" + x)  // <- not allowed for not strings
	xStr := fmt.Sprint(x)

	fmt.Println("1. The value of x is " + xStr)
	fmt.Println("2. The value of x is", x)

	fmt.Printf("3. The value of x is %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "oops"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
