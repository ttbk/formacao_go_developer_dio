package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) imprimirDados() {

	fmt.Println("Nome:", p.nome)
	fmt.Println("Idade:", p.idade)

}

func main() {

	var pes = pessoa{"Thiago", 38}

	pes.imprimirDados()

}
