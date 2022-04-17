package main

import "fmt"

type item struct {
	productId int
	qty       int
	price     float64
}

type purchase struct {
	userId int
	items  []item
}

func (p purchase) totalValue() float64 {
	total := 0.0
	for _, item := range p.items {
		total += item.price * float64(item.qty)
	}
	return total
}

func main() {
	p := purchase{
		userId: 1,
		items: []item{
			{productId: 1, qty: 2, price: 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Total value of purchase is R$ %.2f", p.totalValue())
}
