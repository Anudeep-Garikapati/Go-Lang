package main

import (
	"fmt"
	"go_training/Go-Lang/simplecalc"
)

type Address struct {
	state   string
	pincode int
}

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
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//struct
	vamsi := Address{"HYD", 4584}
	fmt.Println(vamsi)
	Ajay := Address{"AP", 78555}
	fmt.Println(Ajay)

	//arrays & slices
	a1 := [4]int{23, 55, 44, 88}
	fmt.Println(a1)
	fmt.Println(len(a1))
	fmt.Println(cap(a1))
	s1 := []string{"HI", "HELLO", "ALL", "WELCOME", "TO", "THE", "CLASS"}
	fmt.Println(s1[2:5])
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := make([]int, 3, 6)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	s1 = append(s1, "GOLANG")
	fmt.Println(s1)

	//Maps
	m1 := map[int]string{
		1: "where",
		2: "Are",
		3: "you",
	}
	fmt.Println(m1)
	m1[4] = "smile"
	fmt.Println(m1)
}
