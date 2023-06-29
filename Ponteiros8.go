package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr = 42 // Modifica o valor apontado pelo ponteiro
}

func main() {
	value := 10

	fmt.Println("Valor antes da modificação:", value)
	modifyValue(&value)
	fmt.Println("Valor após a modificação:", value)
}
