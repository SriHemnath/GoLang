package main

import "fmt"

const (
	a = iota + 1
	b
	c
)

const (
	d, e, f = iota, iota, iota
)

func main() {
	fmt.Println(a, b, c, d, e, f) // 1 2 3 0 0 0
	fmt.Println("IOTA starts from 0 and increment value after each specification. Values are reset in new declaration")

	//we range over a collection of values using for range
	s := []int{1, 2, 3, 4, 5, 6}
	for index, val := range s {
		fmt.Println("Value at index ", index, "is ", val)
	}

	fmt.Println("Printing prime numbers")
	printPrime()
}

//PRINT PRIME
func printPrime() {
	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
