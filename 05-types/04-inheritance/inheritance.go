package main

import "fmt"

type car struct {
	name     string
	velocity int
}

type ferrari struct {
	car   // anonymous fields (composition)
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.velocity = 0
	f.turbo = true
	fmt.Printf("%s, %v\n", f.name, f.turbo)
	fmt.Println(f)
}
