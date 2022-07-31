package main

import "fmt"

type Printable interface {
	toString() string
}

type Person struct {
	name     string
	lastname string
}

type Product struct {
	name  string
	price float64
}

// Interfaces are implemented inplictly
func (p Person) toString() string {
	return p.name + " " + p.lastname
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x Printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing Printable = Person{"Fausto", "Silva"}
	fmt.Println(thing.toString())
	print(thing)

	thing = Product{"Jeans", 79.90}
	fmt.Println(thing.toString())
	print(thing)

	p2 := Product{"TV", 1499.90}
	fmt.Println(p2.toString())
	print(p2)
}
