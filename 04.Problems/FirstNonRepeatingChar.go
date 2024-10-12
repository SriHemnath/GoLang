package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := make(map[int]string)
	n[1] = "Hemnath"
	n[2] = "Dharani"
	n[4] = "Kalai"
	n[3] = "Arun"
	n[4] = "Bugs"
	fmt.Printf("n: %v\n", n)
	for k, v := range n {
		fmt.Println(k, v)
	}

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["a"] = 7
	m["e"] = 4
	m["c"] = 3

	for k, v := range m {
		fmt.Println(k, v)
	}
	//map does not main insertion order

	a := "hemnath"
	for i, v := range a {
		fmt.Printf("%d %c", i, v)
	}

	fmt.Println(strings.Index(a, "h"))    //will print the 1st index of h
	fmt.Println(findFirst("aaabcccdeef")) //returns 1st non repeating character using map
	fmt.Println(findIndex("aaabcccdeef")) //return 1st non repeating character using indexof
}

func findFirst(str string) map[string]int {
	count := make(map[string]int)
	for _, v := range str {
		count[string(v)] = count[string(v)] + 1
	}
	return count
}

func findIndex(str string) string {
	k := "Hemnath"
	fmt.Println(k[0]) //will print the ascii equivalent of H - 72
	for _, v := range str {
		fmt.Println(v)
		if (strings.LastIndex(str, string(v))) == (strings.Index(str, string(v))) {
			return strconv.QuoteRune(v)
		}
	}
	return "-"
}
