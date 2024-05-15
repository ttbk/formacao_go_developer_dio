package main

import "fmt"

const tempEbulicaoK = 373.0

func main() {

	tempEmK := tempEbulicaoK

	tempEmC := tempEmK - 273.0

	fmt.Printf("A temperatura de ebulição da água em °K é %g e em °C é %g.", tempEmK, tempEmC)

}
