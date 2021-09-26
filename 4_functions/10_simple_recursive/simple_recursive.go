package main

import "fmt"

func factorial(n uint) uint {
	switch {
	case n == 0:
		return 1

	default:
		return n * factorial(n-1)
	}
}

func main() {
	result1 := factorial(5)
	fmt.Println("Fatorial 5 =", result1)

	result2 := factorial(20)
	fmt.Println("Fatorial 20 =", result2)

	// invalid input
	/*_ := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}*/
}
