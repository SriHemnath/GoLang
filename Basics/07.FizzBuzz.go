package main

import "fmt"

func main() {
	// var b int // TODO
	// fmt.Println("Enter limit: ")
	// fmt.Scan(&b)
	// for i := 1; i <= b; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FizzBuzz")
	// 		} else {
	// 			fmt.Println("Fizz")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
	fizzBuzz(15)
}

func fizzBuzz(n int32) {
	// Write your code here

	var i int32 = 1
	for i <= n {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

		i++
	}

}
