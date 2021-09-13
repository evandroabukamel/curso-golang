package main

import "fmt"

func main() {
	// Showing that Go can detect code errors.
	fmt.Printf("Outro programa em %s!")

	/*
		godoc -http=:6060 -> publishes an offline manual at localhost:6060.
		go doc cmd/vet -> show doc for a command.
		go vet file.go -> detect dead code or errors on code.
		go build file.go -> Build an executable.
		go run file.go -> Run a Go program.
	*/
}
