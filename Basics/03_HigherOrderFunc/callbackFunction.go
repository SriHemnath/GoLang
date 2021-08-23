package main

import "fmt"

func filter(number []int, callback func(int) bool) []int {
	var result []int
	for _, v := range number {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	result := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 1
	})
	fmt.Println(result)
}
