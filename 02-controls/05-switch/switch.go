package main

import "fmt"

func gradeToConcept(n float64) string {
	var grade = int(n)
	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 0, 1, 2:
		return "E"
	default:
		return "Invalid grade"
	}
}

func main() {
	fmt.Println(gradeToConcept(8.0))
	fmt.Println(gradeToConcept(10.1))
}
