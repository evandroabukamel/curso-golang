package main

import "fmt"

/*
THERE IS NO TERNARY OPERATOR ON GOLANG.
*/
func getResult(score float64) string {
	if score >= 6 {
		return "Approved"
	}
	return "Repproved"

	// Inexistent alternative
	// return score >= 6 ? "Apporved" : "Repproved"
}

func main() {
	fmt.Println(getResult(5))
}
