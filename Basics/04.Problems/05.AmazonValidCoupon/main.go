package main

import "fmt"

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

func (s *Stack) Peek() string {
	if s.isEmpty() {
		return ""
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element
}

func main() {
	// var first string
	// fmt.Scanln(&first)
	fmt.Println(isValid(""))
	fmt.Println(isValid("vv"))
	fmt.Println(isValid("xbbx"))
	fmt.Println(isValid("bbccdd"))
	fmt.Println(isValid("xyffyxdd"))

	//below are invalid
	fmt.Println(isValid("xyx"))
	fmt.Println(isValid("bbb"))
	fmt.Println(isValid("abc"))
}

func isValid(s string) bool {
	var stack Stack
	for i := 0; i < len(s); i++ {
		c := fmt.Sprintf("%c", s[i])
		if stack.isEmpty() {
			stack.Push(c)
		} else if stack.Peek() == c {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	return stack.isEmpty()
}
