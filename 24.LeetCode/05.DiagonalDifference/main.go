package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 5, 9},
	}
	fmt.Println(diagonalDifference(a)) // (1+5+9) - (3+5+9) = 15 - 17 = 2
	b := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDifference(b)) // ( 11 + 5 + -12) - (4 + 5 + 10) = 4 - 19 = 15
}

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {

	var ldig, rdig int32
	n := len(arr) - 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i == j {
				ldig = ldig + arr[i][j]
			}
			if j == n-i {
				rdig = rdig + arr[i][j]
			}
		}
	}
	return int32(math.Abs(float64(ldig - rdig)))
}
