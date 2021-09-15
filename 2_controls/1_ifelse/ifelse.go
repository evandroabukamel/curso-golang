package main

import "fmt"

func printResult(score float64) {
	if score >= 7 {
		fmt.Println("Approved with score", score)
	} else {
		fmt.Println("Repproved with score", score)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
