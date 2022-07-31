package main

import "fmt"

type Course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type dynamic interface{}

	var other dynamic
	fmt.Println(other)

	other = true
	fmt.Println(other)

	other = Course{name: "Golang course"}
	fmt.Println(other)

}
