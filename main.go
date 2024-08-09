package main

import (
	"fmt"
	"go_training/simplecalc"
)

func main() {
	fmt.Println("Hell")
	a, b := 11.0, 9.0
	fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Sub(a, b))
}
