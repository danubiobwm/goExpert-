package main

import (
	"fmt"
)

func main() {

	fmt.Println(soma(1, 5, 45, 100, 545, 600))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
