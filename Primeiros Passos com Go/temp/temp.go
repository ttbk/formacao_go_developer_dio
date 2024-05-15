package main

import "fmt"

const ebulicaoF = 212.0

func main() {

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("Temp de ebulição da água em °F é: %f e Temp de ebulição da água em °C é: %F", tempF, tempC)

}
