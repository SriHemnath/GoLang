package main

import "fmt"

func main() {
	fmt.Println(fib(0, 1, 10))
}

func fib(a, b, upTo int) int {
	out := make(chan int)
	go fibR(a, b, upTo, out)
	fmt.Println("After calling go routine")
	return <-out
}

func fibR(a, b, upTo int, ch chan<- int) {
	fmt.Println("goroutine fibR running")
	c := a + b
	if c >= upTo {
		ch <- c
		fmt.Println("Returning conrol to main")
		return
	}
	go fibR(b, c, upTo, ch)
}
