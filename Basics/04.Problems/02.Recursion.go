package main

import "fmt"

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}

func main() {
	fmt.Println(recursion(5))
}
