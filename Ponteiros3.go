package main

import (
	"fmt"
	"unsafe"
)

func reverseString(strPtr *string) {
	// Converter o ponteiro para um slice de bytes mutável
	strBytes := []byte(*strPtr)

	// Inverter o slice de bytes
	for i, j := 0, len(strBytes)-1; i < j; i, j = i+1, j-1 {
		strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	}

	// Converter o slice de bytes de volta para uma string
	*strPtr = *(*string)(unsafe.Pointer(&strBytes))
}

func main() {
	str := "Hello, World!"
	fmt.Println("Antes:", str)

	// Passar o endereço da string para a função reverseString
	reverseString(&str)

	fmt.Println("Depois:", str)
}
