package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(l *Livro) {
	l.Preco = l.Preco * 0.9
}

func main() {
	livro := Livro{
		Titulo: "A Guerra dos Tronos",
		Autor:  "George R. R. Martin",
		Preco:  50.0,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)
	aplicarDesconto(&livro)
	fmt.Println("Preço após o desconto:", livro.Preco)
}
