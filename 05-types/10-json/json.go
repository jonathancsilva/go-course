package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{1, "Notebook", 1899.90, []string{"Promotion", "New"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct
	var p2 product
	jsonString := `{"id":2,"name":"Cellphone","price":499.9,"tags":["Promotion","New"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
