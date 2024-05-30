package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	SobreNome string    `json:"sobrenome"`
	Endereco  *Endereco `json:"endereco"`
}

type Endereco struct {
	Rua    string `json:"rua"`
	Numero string `json:"numero"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
}

var people []Person

func GetPersonList(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(people)
}

func GetPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(res).Encode(person)
			return
		}
	}
	json.NewEncoder(res).Encode(&Person{})
}

func PostPerson(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(res).Encode(people)

}

func DeletePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(res).Encode(people)
	}

}

func main() {

	people = append(people, Person{ID: "1", Nome: "Thiago", SobreNome: "Cesar", Endereco: &Endereco{Rua: "Nova", Numero: "123", Cidade: "Olinda", Estado: "PE"}})
	people = append(people, Person{ID: "2", Nome: "Wal", SobreNome: "Cesar", Endereco: &Endereco{Rua: "Antiga", Numero: "321", Cidade: "Recife", Estado: "PE"}})

	router := mux.NewRouter()
	router.HandleFunc("/person", GetPersonList)
	router.HandleFunc("/person/{id}", GetPerson)
	router.HandleFunc("/person/{id}", PostPerson)
	router.HandleFunc("/person/{id}", DeletePerson)
	log.Fatal(http.ListenAndServe(":8000", router))

}
