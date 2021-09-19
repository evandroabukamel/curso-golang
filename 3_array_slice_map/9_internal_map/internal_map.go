package main

import "fmt"

func main() {
	employees := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"Jose Joao": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	fmt.Println(employees)

	delete(employees, "P")
	fmt.Println(employees)

	for letter, empls := range employees {
		fmt.Println(letter, empls)
	}
}
