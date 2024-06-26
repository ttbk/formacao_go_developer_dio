package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	site := "https://gobyexample.com"

	resp, err := http.Get(site)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Qual o statu:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
