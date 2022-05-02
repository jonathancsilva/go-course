package main

import (
	"fmt"
	"time"
)

func multiple(number int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * number

	time.Sleep(time.Second)
	c <- 3 * number

	time.Sleep(3 * time.Second)
	c <- 4 * number
}

func main() {
	c := make(chan int)
	go multiple(2, c)

	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println(<-c)
}
