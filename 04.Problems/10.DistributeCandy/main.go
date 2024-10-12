package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 1}
	fmt.Println(totalCandy(arr))
	fmt.Println(totalCandy([]int{1, 2, 2}))
	fmt.Println(totalCandy([]int{1, 3, 2, 1}))
	fmt.Println(totalCandy([]int{1, 2}))
}

func totalCandy(arr []int) int {
	total := len(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			total++
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			total++
		}
	}

	return total
}
