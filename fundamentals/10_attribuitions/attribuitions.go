package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	// Declaration with attribuitions with type inference
	i := 3
	i += 3 // i = 1 + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	// Multiple declaration
	x, y := 1, 2
	fmt.Println(x, y)

	// Values swap
	x, y = y, x
	fmt.Println(x, y)
}
