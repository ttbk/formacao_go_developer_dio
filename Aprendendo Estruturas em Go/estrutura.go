package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {

	p := pessoa{"Thiago", 38}

	fmt.Println(p)
	fmt.Println(p.nome)
	fmt.Println(p.idade)

}
