package main

import (
	"fmt"
	"sort"
)

//a = [1,2,3,4,2,3,4,5,3,3,5,6]
//op = [3 2 4 5 1 6]

func main() {
	sl := []int{1, 2, 3, 4, 2, 3, 4, 5, 3, 3, 5, 6}
	unique := make(map[int]int)
	keys := make([]int, 0, len(sl))
	for _, value := range sl {
		if val, ok := unique[value]; ok {
			unique[value] = val + 1
		} else {
			keys = append(keys, value)
			unique[value] = 1
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		// return unique[keys[i]] > unique[keys[j]]
		return i < j

	})

	fmt.Println(unique)
	fmt.Println(keys)
}
