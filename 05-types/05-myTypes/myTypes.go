package main

import "fmt"

type grade float64

func (g grade) between(init, ended float64) bool {
	return float64(g) >= init && float64(g) <= ended
}

func returnGrade(g grade) string {
	if g.between(9.0, 10.0) {
		return "A"
	} else if g.between(7.0, 8.99) {
		return "B"
	} else if g.between(5.0, 7.99) {
		return "C"
	} else if g.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(returnGrade(9.8))
	fmt.Println(returnGrade(6.9))
	fmt.Println(returnGrade(2.1))
}
