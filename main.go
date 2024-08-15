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

	day := 2

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
