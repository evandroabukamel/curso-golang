package main

import "fmt"

func main() {
	// homogeneous (same type) and staic (fixed size)
	var scores [3]float64
	fmt.Println(scores)

	scores[0], scores[1], scores[2] = 7.8, 4.3, 9.1
	// scores[3] = 10  // invalid array index 3 (out of bounds for 3-element array)
	fmt.Println(scores)

	total := 0.0
	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}

	avg := total / float64(len(scores))
	fmt.Printf("Average score %.2f\n", avg)
}
