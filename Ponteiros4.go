package main

import "fmt"

func sumLastDigits(numPtr *int) {
	// Obter o valor da variável apontada pelo ponteiro
	num := *numPtr

	// Extrair os dois últimos dígitos
	lastDigit := num % 10
	num /= 10
	secondLastDigit := num % 10

	// Calcular a soma dos dois últimos dígitos
	sum := lastDigit + secondLastDigit

	// Atualizar o valor da variável para a soma
	*numPtr = sum
}

func main() {
	// Exemplo de uso da função
	num := 1234
	fmt.Println("Valor inicial:", num)

	sumLastDigits(&num)
	fmt.Println("Novo valor:", num)
}
