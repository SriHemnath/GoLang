package main

import "fmt"

func main() {
	b := 1
	x := []int{1, 2, 3, 4, 5, 6}
	//This won't compile as go doesn't allow postincrement to be used as expression
	//They can be used only as statements
	// a := 5*4 + x[b++] - (9 / b)  In java ans will be 18
	b++ //This is allowed,  *preincrement is not allowed anywhere
	a := 5*4 + x[b] - (9 / b)
	fmt.Println(a)
}
