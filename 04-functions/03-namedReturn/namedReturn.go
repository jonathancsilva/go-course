package main

import "fmt"

func change(p1, p2 int) (second, frist int) {
	frist = p1
	second = p2
	return // clean return
}

func main() {
	x, y := change(1, 2)
	fmt.Println(x, y)

	second, frist := change(1, 9)
	fmt.Println(second, frist)
}
