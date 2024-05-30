//http://pokeapi.co/api/v2/pokedex/kanto

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// type Descriptions struct{
// 	descricoes []Description `json:"descriptions"`
// }

// type Description struct{
// 	descricao  string `json:"description"`
// 	linguagem Linguagem `json:"language"`
// }

// type Linguagem struct {
// 	nome string `json:"name"`
// 	url string `json:"url"`
// }

type Pokemons struct {
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Entry_number    int     `json:"entry_number"`
	Pokemon_species Species `json:"pokemon_species"`
}

type Species struct {
	Nomee string `json:"name"`
	Url   string `json:"url"`
}

func main() {

	apiSite := "http://pokeapi.co/api/v2/pokedex/kanto"

	resp, err := http.Get(apiSite)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println((string(responseData)))

	var responseObject Pokemons

	json.Unmarshal(responseData, &responseObject)

	fmt.Sprint("Tamanho:", string(len(responseObject.Pokemon)))

	fmt.Println("Numero: ", responseObject.Pokemon[0].Entry_number)
	fmt.Println("Nome: ", responseObject.Pokemon[0].Pokemon_species.Nomee)
	fmt.Println("URL: ", responseObject.Pokemon[0].Pokemon_species.Url)

}
