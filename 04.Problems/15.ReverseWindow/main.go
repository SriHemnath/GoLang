package main

import "fmt"

//
//sl := []int[1,2,3,4,5,6,7,8]
//window:= 3
//op := []int{3,2,16,5,4,7}
func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	window := 3
	fmt.Println(reverseWindow(window, sl))
	fmt.Println(reverseWindow(4, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func reverseWindow(w int, value []int) []int {
	result := make([]int, 0)
	lastIndex := 0
	for i := 0; i+w < len(value); i = i + w {
		sl := value[i : i+w]
		reverse(sl)
		result = append(result, sl...)
		lastIndex = i + w
	}
	sl := value[lastIndex:]
	reverse(sl)
	result = append(result, sl...)
	return result
}

func reverse(value []int) {
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}

}
