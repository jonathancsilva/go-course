package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)

	// bitwise
	fmt.Println(a & b) // 11 & 10 = 10
	fmt.Println(a | b) // 11 | 10 = 11
	fmt.Println(a ^ b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println(math.Max(float64(a), float64(b)))
	fmt.Println(math.Min(c, d))
	fmt.Println(math.Pow(c, d))
}
