package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	note := 6.9
	finalNote := int(note)
	fmt.Println(finalNote)

	// Careful...
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123") // returns => num, error
	fmt.Println("Test ", num-122)

	// boolean, true == 1
	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("b is true")
	} else {
		fmt.Println("b is false")
	}
}
