package main

import "fmt"

func main() {
	employees := map[string]map[string]float64{
		"G": {
			"Gabriela": 9893.33,
			"Gustavo":  9833.33,
		},
		"J": {
			"João": 939393.30,
		},
		"P": {
			"Pedro": 636844.30,
		},
	}

	fmt.Println(employees)
	delete(employees, "P")
	fmt.Println(employees)

	for letter, employee := range employees {
		fmt.Println(letter, employee)

		for name, salary := range employee {
			fmt.Println(name, salary)
		}
	}
}
