package main

import "fmt"

func getResult(grade float64) string {
	if grade >= 6 {
		return "Approved"
	}
	return "Reproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
