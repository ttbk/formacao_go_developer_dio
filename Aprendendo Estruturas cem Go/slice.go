package main

import (
	"fmt"
)

func main() {
	a := [3]int{0, 0, 0}
	v := a[:]
	fmt.Printf("len: %v cap: %v\n", len(v), cap(v))
	for i := 0; i <= 4; i++ {
		v = append(v, i)
		fmt.Printf("len: %v cap: %v\n", len(v), cap(v))
		fmt.Println(v)
	}
	fmt.Println(v)
}
