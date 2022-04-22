package main

import "fmt"

type sporting interface {
	turboOn()
}

type luxurious interface {
	park()
}

type sportingLuxurious interface {
	sporting
	luxurious
}

type bmw7 struct{}

func (b bmw7) turboOn() {
	fmt.Println("turboOn...")
}

func (b bmw7) park() {
	fmt.Println("park...")
}

func main() {
	var b sportingLuxurious = bmw7{}
	b.turboOn()
	b.park()

	fmt.Println(b)
}
