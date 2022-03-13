package main

import "fmt"

func purchases(job1, job2 bool) (bool, bool, bool) {
	buyTv50 := job1 && job2
	buyTv32 := job1 != job2 // xor
	buyIcecream := job1 || job2

	return buyTv50, buyTv32, buyIcecream
}

func main() {
	tv50, tv32, icecream := purchases(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, icecream, !icecream)
}
