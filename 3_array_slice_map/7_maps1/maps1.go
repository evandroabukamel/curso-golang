package main

import "fmt"

func main() {
	// Maps must be inicializaed
	// var approved map[string]string
	approved := make(map[string]string)

	approved["000111222333"] = "Joao"
	approved["123456678990"] = "Maria"
	approved["098765432132"] = "Jesus"
	fmt.Println(approved)

	for ssn, name := range approved {
		fmt.Printf("%s (SSN: %s)\n", name, ssn)
	}

	fmt.Println(approved["123456678990"])
	delete(approved, "123456678990")
	fmt.Println(approved["123456678990"])
}
