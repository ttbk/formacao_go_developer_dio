package main

import "fmt"

func main() {

	numIni := 1
	numFin := 100

	for i := numIni; i <= numFin; i++ {

		if i%3 == 0 && i%5 == 0 {

			fmt.Println("Pin! Pan!")

		} else if i%3 == 0 {

			fmt.Println("Pin!")

		} else if i%5 == 0 {

			fmt.Println("Pan!")

		} else {

			fmt.Println(i)

		}

	}

}
