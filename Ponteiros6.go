package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func AlterarTituloSeAutorAnonimo(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro Exemplo",
		Autor:  "Anônimo",
	}

	fmt.Println("Antes da alteração:", livro.Titulo)
	AlterarTituloSeAutorAnonimo(&livro)
	fmt.Println("Após a alteração:", livro.Titulo)
}
