package main

import "fmt"

func scoreToConcept(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(scoreToConcept(9.8))
	fmt.Println(scoreToConcept(4.3))
	fmt.Println(scoreToConcept(2.5))
}
