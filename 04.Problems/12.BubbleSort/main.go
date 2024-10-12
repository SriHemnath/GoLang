package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 10, 4, 5, 6}
	for i := 0; i < len(arr); i++ {
		swap := 0 //if the swap is not incremented then array is in sorted order
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swap++
			}
		}
		if swap == 0 {
			break
		}
	}
	fmt.Println(arr)
}
