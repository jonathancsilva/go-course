package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getName() string {
	return p.name + " " + p.lastname
}

func (p *person) setName(name string) {
	split := strings.Split(name, " ")
	p.name = split[0]
	p.lastname = split[1]
}

func main() {
	p := person{"Jonathan", "Silva"}
	fmt.Println(p.getName())

	p.setName("Gustavo Silva")
	fmt.Println(p.getName())
}
