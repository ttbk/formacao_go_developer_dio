package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"usuarios"`
}

type User struct {
	Nome   string `json:"Nome"`
	Tipo   string `json:"Tipo"`
	Idade  int    `json:"Idade"`
	Social Social `json:"Social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("usuarios.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Arquivo Aberto com Sucesso")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var usuarios Users

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Users); i++ {
		fmt.Println("Usuário Tipo: " + usuarios.Users[i].Tipo)
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Users[i].Idade))
		fmt.Println("Usuário Nome: " + usuarios.Users[i].Nome)
		fmt.Println("Usuário Social F: " + usuarios.Users[i].Social.Facebook)
		fmt.Println("Usuário Social T: " + usuarios.Users[i].Social.Twitter)

	}

}
