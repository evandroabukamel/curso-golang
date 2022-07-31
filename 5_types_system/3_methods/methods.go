package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name     string
	lastname string
}

func (p Person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *Person) setFullName(fullname string) {
	parts := strings.Split(fullname, " ")
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := Person{"Evandro", "Lopes"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Jose Antonio")
	fmt.Println(p1.getFullName())
}
