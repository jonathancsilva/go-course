package main

import "fmt"

func main() {
	// var approved map[int]string
	approved := make(map[int]string)

	approved[1234567890] = "Maria"
	approved[9849384939] = "Pedro"
	approved[7674848487] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[9849384939])
	delete(approved, 9849384939)
	fmt.Println(approved)
}
