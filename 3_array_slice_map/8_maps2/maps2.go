package main

import "fmt"

func main() {
	salaries := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}
	fmt.Println(salaries)

	salaries["Rafael Filho"] = 1350.0
	fmt.Println(salaries)
	delete(salaries, "unknown") // No error

	for name, salary := range salaries {
		fmt.Println(name, salary)
	}
}
