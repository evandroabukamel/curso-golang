package main

import "fmt"

type Item struct {
	productID int
	quantity  int
	price     float64
}

type Order struct {
	userID int
	items  []Item
}

func (o Order) getTotalPrice() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := Order{
		userID: 1,
		items: []Item{
			Item{1, 2, 12.10},
			Item{2, 1, 23.49},
			Item{11, 100, 3.13},
		},
	}

	fmt.Printf("Order total price: %.2f", order.getTotalPrice())
}
