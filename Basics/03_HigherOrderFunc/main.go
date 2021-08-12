package main

import "fmt"

func main() {

	//partial holds the returned function
	partial := partialSum(3)
	fmt.Println(partial(2))

}

//takes int as parameter and returns a func
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func sum(x, y int) int {
	return x + y
}
