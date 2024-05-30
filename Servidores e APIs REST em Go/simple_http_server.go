package main

import (
	"fmt"
	"net/http"
)

func ola(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Ola\n")
}

func cabecalhos(res http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {

		fmt.Fprintf(res, "TESTE CABEÇALHO: ")

		for _, c := range cabecalhos {
			fmt.Fprintf(res, "%v: %v\n", nome, c)
		}

	}

}

func main() {

	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)
	http.ListenAndServe(":8090", nil)

}
