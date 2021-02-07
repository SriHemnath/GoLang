package main

import (
	"fmt"
	"os"
)

func main() {
	s1 := "Hemnath"
	s2 := "Hemnath"
	s3 := &s1       //gets memory loc of s1 and assigns it to s3. Retrive value in s3 by using *

	println("Address of s1 ",&s1)
	println("Address of s2 ",&s2)
	println("Address of s3 ",&s3)
	println("Not like Java. Every time value is stored in new memory locaiton.")
	println("We can handle memory efficiently by using pointers")
	println(s1 == s2)
	println(s1 == *s3)

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("arguments 0 will always be currently running executable ",os.Args[0])
	fmt.Println("It's over", os.Args[1])
}