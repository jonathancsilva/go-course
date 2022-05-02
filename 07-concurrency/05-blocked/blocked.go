package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // block operation
	fmt.Println("only after read this")
}

func main() {
	c := make(chan int) // channel without buffer
	go routine(c)

	fmt.Println(<-c) // block operation
	fmt.Println("read..")
	fmt.Println(<-c)
	fmt.Println("end..")
}
