package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		} else {
			fmt.Println("Odd", i)
		}
	}

	for {
		// infinite loop
		fmt.Println("Forever..")
		time.Sleep(time.Second)
	}
}
