package main

import (
	"fmt"
	m "math" // m is a reference
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // type (float64)

	area := PI * m.Pow(raio, 2)
	fmt.Println("The area of the circle is", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "ops!"
	fmt.Println(g, h, i)
}
