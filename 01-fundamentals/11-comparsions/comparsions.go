package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Banana" == "Banana")
	fmt.Println(3 != 2)
	fmt.Println(3 < 2)
	fmt.Println(3 > 2)
	fmt.Println(3 <= 2)
	fmt.Println(3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Person struct {
		Name string
	}

	p1 := Person{"João"}
	p2 := Person{"João"}
	fmt.Println(p1 == p2)
}
