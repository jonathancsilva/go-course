package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)        // by value
	fmt.Println(n) // not change

	inc2(&n)       // by reference
	fmt.Println(n) // increment value
}
