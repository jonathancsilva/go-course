package main

import "fmt"

type printing interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(p printing) {
	fmt.Println(p.toString())
}

func main() {
	var p1 printing = person{
		name:     "Jonathan",
		lastname: "Silva",
	}
	fmt.Println(p1.toString())
	print(p1)

	p2 := product{
		name:  "Computer",
		price: 999.00,
	}
	fmt.Println(p2.toString())
	print(p2)
	print(product{"T-Shirt", 45.99})
}
