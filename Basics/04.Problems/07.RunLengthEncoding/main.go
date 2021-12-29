package main

import (
	"fmt"
)

func main() {
	RunLengthEncoding("aaabbccc")
	RunLengthEncoding("aaabbcca")
}

func RunLengthEncoding(raw string) {
	if len(raw) == 0 {
		fmt.Println("-")
		return
	}
	current := 0
	nextPos := 1
	count := 1
	for {
		if raw[current] == raw[nextPos] {
			count++
			current++
			nextPos++
		} else {
			fmt.Print(string(raw[current]), count)
			count = 1
			current++
			nextPos++
		}
		if len(raw) == nextPos {
			fmt.Print(string(raw[current]), count)
			fmt.Println("")
			break
		}
	}
}
