package main

import (
	"fmt"
)

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func main() {
	var stack Stack
	stack.Push("Hemnath")
	stack.Push("Alia")
	stack.Push("Kajal")
	stack.Push("Samantha")
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	for !stack.isEmpty() {
		fmt.Println(stack.Pop())
	}
	fmt.Println(stack.isEmpty())
}
