package main

import "fmt"

func main() {

	numIni := 1
	numFin := 100

	for i := numIni; i <= numFin; i++ {

		if i%3 == 0 {

			fmt.Println(i)

		}

	}

}
