package main

import "fmt"

func main() {
	ratings := []int{2, 1, 3}
	fmt.Println(countDecreasingRating(ratings))              //4
	fmt.Println(countDecreasingRating([]int{4, 3, 5, 4, 3})) //9
	fmt.Println(countDecreasingRating([]int{3, 2, 1}))       //6
	fmt.Println(countDecreasingRating([]int{1}))             //1
	fmt.Println(countDecreasingRating([]int{9, 8, 7, 6, 5})) //15
}

func countDecreasingRating(ratings []int) int {
	load, ans := 0, 0
	for i := 0; i < len(ratings); i++ {
		if load == 0 {
			load = 1
		} else {
			if ratings[i-1] == ratings[i]+1 {
				load++
			} else {
				ans += sumOfNNaturalNumbers(load)
				load = 1
			}
		}
	}
	ans += sumOfNNaturalNumbers(load)
	return ans
}

func sumOfNNaturalNumbers(n int) int {
	if n%2 == 0 {
		return (n / 2) * (n + 1)
	}
	return n * ((n + 1) / 2)
}
