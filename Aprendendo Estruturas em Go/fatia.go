package main

import "fmt"

func main() {

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 1)

	fmt.Println(fatia1, fatia2)

	copy(fatia2, fatia1)

	fmt.Println(fatia1, fatia2)

}
