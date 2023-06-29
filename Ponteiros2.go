package main

import "fmt"

func checkParImpar(valor *int) {
	if *valor%2 == 0 {
		*valor = 0 // inteiro é par
	} else {
		*valor = 1 // inteiro é ímpar
	}
}

func main() {
	numero := 7
	fmt.Println("Antes:", numero)

	checkParImpar(&numero)
	fmt.Println("Depois:", numero)
}
