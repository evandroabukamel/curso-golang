package main

import "fmt"

func printApproved(approveds ...string) {
	fmt.Println("Approved list")

	for i, name := range approveds {
		fmt.Printf("%d) %s\n", i+1, name)
	}
}

func main() {
	approveds := []string{"Mary", "Peter", "Raphael", "Donatelo"} // slice
	printApproved(approveds...)
}
