package main

import "fmt"

/*
THERE IS NO TERNARY OPERATOR ON GOLANG.
*/
func getResult(note float64) string {
	if note >= 6 {
		return "Approved"
	}
	return "Repproved"

	// Inexistent alternative
	// return note >= 6 ? "Apporved" : "Repproved"
}

func main() {
	fmt.Println(getResult(5))
}
