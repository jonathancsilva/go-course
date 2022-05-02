package main

import (
	"fmt"

	"github.com/jonathancsilva/html"
)

func forward(origem <-chan string, destiny chan string) {
	for {
		destiny <- <-origem
	}
}

func union(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(entry1, c)
	go forward(entry2, c)
	return c
}

func main() {
	c := union(
		html.Title("https://www.cod3r.com.br", "https://www.google.com"),
		html.Title("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c)
}
