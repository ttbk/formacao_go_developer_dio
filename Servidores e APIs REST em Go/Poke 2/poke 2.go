package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Pokemons struct {
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Entry_number    int     `json:"entry_number"`
	Pokemon_species Species `json:"pokemon_species"`
}

type Species struct {
	Nome string `json:"poke_name"`
	Url  string `json:"poke_url"`
}

func main() {

	jsonFile, err := os.Open("pokemonApi.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Arquivo Aberto com Sucesso")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var pokes Pokemons

	json.Unmarshal(byteValue, &pokes)

	fmt.Println("TAMANHO:", len(pokes.Pokemon))

	for i := 0; i < len(pokes.Pokemon); i++ {
		fmt.Println("I =", i)
		fmt.Println("Numero: " + strconv.Itoa(pokes.Pokemon[i].Entry_number))
		fmt.Println("Nome: " + pokes.Pokemon[i].Pokemon_species.Nome)
		fmt.Println("URL: " + pokes.Pokemon[i].Pokemon_species.Url)

	}

}
