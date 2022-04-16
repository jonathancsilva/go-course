package main

import "fmt"

func getValueApproved(number int) int {
	defer fmt.Println("end!")
	if number >= 5000 {
		fmt.Println("High number")
		return number
	}
	fmt.Println("Low number")
	return number
}

func main() {
	fmt.Println(getValueApproved(6000))
	fmt.Println(getValueApproved(4000))
}
