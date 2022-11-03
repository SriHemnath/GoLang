package main

import "fmt"

/*
size n=3
2 -> array of 2 elements
[[0,2],[1,3],[0,7]]
elements represent characters from a-z in 0 to 25
output : a
2-0 = 2 for a
3-2 = 1 for b
worst case 7-3=4 for a
*/

func main() {
	alphabets := []string{"a", "b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z"}
	keys := [][]int{{0, 2}, {1, 3}, {0, 7}}
	var slowest int
	var slowestChar int
	var previous = 0
	for _, key := range keys {
		processTime := key[1] - previous
		if processTime < slowest {
			slowest = processTime
			slowestChar = key[0]
		}
	}
	fmt.Println(alphabets[slowestChar])
}
