package main

import (
	"fmt"
	"time"
)

func types(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Not recognize"
	}
}

func main() {
	fmt.Println(types(2.3))
	fmt.Println(types(1))
	fmt.Println(types("abc"))
	fmt.Println(types(func() {}))
	fmt.Println(types(true))
	fmt.Println(types(time.Now()))
}
