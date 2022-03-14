package main

import "fmt"

func main() {
	employees := map[string]float64{
		"João":  11345.33,
		"Maria": 15349.99,
		"Pedro": 8444.54,
	}

	employees["Rafael"] = 1350.00
	delete(employees, "Gustavo")

	for name, salary := range employees {
		fmt.Println(name, salary)
	}
}
