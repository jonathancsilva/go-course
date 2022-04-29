package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qty int) {
	for i := 0; i < qty; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// speak("Maria", "Por que você não fala comigo?", 3)
	// speak("João", "Só posso falar depois de você!", 1)

	// go speak("Maria", "Ei..", 10)
	// go speak("João", "Opa..", 10)

	go speak("Maria", "Entendi!!", 10)
	speak("João", "Parabéns!!", 5)
}
