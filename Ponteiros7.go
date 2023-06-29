package main

import "fmt"

type Conta struct {
	saldo   float64
	titular string
}

func adicionarSaldo(conta *Conta, valor float64) {
	conta.saldo += valor
}

func main() {
	// Exemplo de uso da função
	contaExemplo := Conta{saldo: 100.0, titular: "João"}
	fmt.Println("Saldo antes:", contaExemplo.saldo)

	adicionarSaldo(&contaExemplo, 50.0)
	fmt.Println("Saldo depois:", contaExemplo.saldo)
}
