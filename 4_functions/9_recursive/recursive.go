package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d\n", n)

	case n == 0:
		return 1, nil

	default:
		previous, _ := factorial(n - 1)
		return n * previous, nil
	}
}

func main() {
	result1, _ := factorial(5)
	fmt.Println("Fatorial 5 =", result1)

	result2, _ := factorial(20)
	fmt.Println("Fatorial 20 =", result2)

	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// Better solution would be...
}
