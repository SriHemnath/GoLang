package main

import "fmt"

func main() {

	var stack []string

	stack = append(stack, "world!")
	stack = append(stack, "Hello ")

	for len(stack) > 0 {

		// Print top
		n := len(stack) - 1
		fmt.Print(stack[n])

		// Pop
		stack = stack[:n]

	}

}
