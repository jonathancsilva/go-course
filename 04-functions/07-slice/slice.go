package main

import "fmt"

func printApproved(approved ...string) {
	fmt.Println("List of Approved:")
	for i, approved := range approved {
		fmt.Printf("%d) %s\n", i, approved)
	}
}

func main() {
	approved := []string{"Maria", "Pedro", "Rafael"}
	printApproved(approved...)
}
