package main

import "fmt"

func main() {
	fmt.Println(compareTriplets([]int32{1, 2, 3}, []int32{3, 2, 1}))       //[1,1]
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))      //[1,1]
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 18})) //[2,1]
}

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {

	result := make([]int32, 2)

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			result[0]++
		} else if a[i] < b[i] {
			result[1]++
		}
	}
	return result
}
