package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	result := average(7.7, 8.1, 5.9, 9.9)
	fmt.Printf("Average: %.2f", result)
}
