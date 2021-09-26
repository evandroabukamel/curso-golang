package main

import "fmt"

func inc1(n int) {
	n++
}

// revision: pointer is represented by *
// In this case, * is used to dereference, have access, to the valores where the pointer points.
func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) // by value
	fmt.Println(n)

	// Revision & is used to get a variable address
	inc2(&n) // by reference
	fmt.Println(n)
}
