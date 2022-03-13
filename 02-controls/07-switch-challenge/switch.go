package main

import "fmt"

func gradeToConcept(grade float64) string {
	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade >= 8 && grade < 9:
		return "B"
	case grade >= 5 && grade < 8:
		return "C"
	case grade >= 3 && grade < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(gradeToConcept(9.0))
	fmt.Println(gradeToConcept(8.6))
	fmt.Println(gradeToConcept(7.7))
	fmt.Println(gradeToConcept(3.5))
	fmt.Println(gradeToConcept(2.0))
}
