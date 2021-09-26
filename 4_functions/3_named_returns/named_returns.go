package main

import "fmt"

func swap(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return // clean return, returns the named returns
}

func main() {
	x, y := swap(2, 3)
	fmt.Println(x, y)

	second, first := swap(7, 1)
	fmt.Println(second, first)
}
