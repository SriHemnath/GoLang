package main

import "fmt"

//a has package level scope
var a = 0

//anymodifications done to a will reflect in other places too
//declare it as const to avoid modifications in different places
func increment() int {
	a++
	return a
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(a)
}
