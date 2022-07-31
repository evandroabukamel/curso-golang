package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

// Metrhod: receiver function
func (p Product) getPriceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 Product
	product1 = Product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.05,
	}
	fmt.Println(product1, product1.getPriceWithDiscount())

	product2 := Product{"Notebook", 1989.90, 0.10}
	fmt.Println(product2, product2.getPriceWithDiscount())
}
