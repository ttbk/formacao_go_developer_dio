package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() {

	lista := []float64{98, 99, 77, 82, 50}
	fmt.Println(media(lista))

}
