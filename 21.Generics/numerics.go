package main

import "fmt"

func max[T int | float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(max(7, 5))
}
