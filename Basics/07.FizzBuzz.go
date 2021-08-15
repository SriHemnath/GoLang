package main

import "fmt"

func main() {
	var b int
	fmt.Println("Enter limit: ")
	fmt.Scan(&b)
	for i := 1; i <= b; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
