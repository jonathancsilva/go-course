package main

import "fmt"

type sporting interface {
	turboOn()
}

type ferrari struct {
	model string
	turbo bool
	speed int
}

func (f *ferrari) turboOn() {
	f.turbo = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turboOn()

	var car2 sporting = &ferrari{"F40", false, 0}
	car2.turboOn()

	fmt.Println(car1, car2)
}
