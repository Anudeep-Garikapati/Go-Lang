package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	x, y := <-c1, <-c2
	fmt.Println(x, y, x+y)

	//PRIME NUMBER GENERATOR

	ch := make(chan int)
	generate(ch)

	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

}

func generate(ch chan int) {
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
}

func filter(in chan int, out chan int, filter int) {
	for val := range in {
		if val%filter != 0 {
			out <- val
		}
	}
}
