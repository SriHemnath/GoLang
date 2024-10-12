package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 9, 1, 3, 5, 3, 3}
	fmt.Println(majorityElement(arr, len(arr)))
	fmt.Println(majorityElement([]int{8, 8, 8, 8, 8, 10, 10}, 7))
}

func majorityElement(arr []int, n int) int {
	// fmt.Println("Before sorting: ", arr)
	sort.Ints(arr)
	// fmt.Println("After sorting: ", arr)
	count := 1
	for i := 1; i < n; i++ {
		for i < n && arr[i] == arr[i-1] {
			count++
			i++
		}
		if count > n/2 {
			return arr[i-1]
		}
		count = 1
	}
	return -1
}
