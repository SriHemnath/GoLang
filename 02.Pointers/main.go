package main

import "fmt"

func main() {

	a := 5
	var b *int
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Name: %v, Memory address of name: %v\n", name, &name)
	fmt.Println("Value of a: ", a)
	// Storing the address of a into b which is of type pointer to int
	b = &a

	// b has the address of a. To access value in b use *
	fmt.Println("Printing address of a: ", &a)
	fmt.Println("Printing value in b: ", b) //which will be address of a
	fmt.Println("Checking address of a and value in b is same ", b == &a)
	fmt.Println("Printing address of b: ", &b)
	fmt.Println("Deferencing the memory address in b", *b)
	*b = 6
	fmt.Println("Modifying *b will reflect in a: ", a)
	increment(b)
	fmt.Println(a)

}

func increment(x *int) {
	*x++
}
