package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	score := 6.9
	finalScore := int(score)
	fmt.Println(finalScore)

	// Careful...
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123") // returns => num, error
	fmt.Println("Test ", num-122)

	// boolean, true == 1
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("b is", strconv.FormatBool(b))
	} else {
		fmt.Println("b is", strconv.FormatBool(b))
	}
}
