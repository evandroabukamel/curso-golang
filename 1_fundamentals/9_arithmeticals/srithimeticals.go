package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtract =>", a-b)
	fmt.Println("Divide =>", a/b)
	fmt.Println("Multiply =>", a*b)
	fmt.Println("Mod =>", a%b)

	// bitwise
	fmt.Println("AND =>", a&b) // 11 & 10 == 10
	fmt.Println("OR  =>", a|b) // 11 | 10 = 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	// math operations
	c := 3.0
	d := 2.0

	fmt.Println("Greater a, b =>", math.Max(float64(a), float64(b)))
	fmt.Println("Lesser c, d =>", math.Min(c, d))
	fmt.Println("Power cˆd =>", math.Pow(c, d))
}
