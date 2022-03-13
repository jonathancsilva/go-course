package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println(reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println(i1)
	fmt.Println(reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))

	var x float32 = 49.99
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(49.99))

	bo := true
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Hello, my name is Jonathan"
	fmt.Println(s1 + "!")
	fmt.Println(len(s1))
	fmt.Println(reflect.TypeOf(s1))

	s2 := `Hello
	my
	name
	is
	Jonathan`
	fmt.Println(s2 + "!")
	fmt.Println(len(s2))
	fmt.Println(reflect.TypeOf(s2))

	char := 'a'
	fmt.Println(char)
	fmt.Println(reflect.TypeOf(char))
}
