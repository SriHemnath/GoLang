package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s1 := "Hemnath"
	s2 := "Hemnath"
	s3 := &s1 //gets memory loc of s1 and assigns to s3. Access value in s3 by using `*`
	s4 := s1[:3]
	fmt.Println(s4, s4[2])

	println("Address of s1 ", &s1)
	println("Address of s2 ", &s2)
	println("Address of s3 ", &s3)
	println("Unlike java go does not refer to the same var location if duplicate value is created")
	println("s1 and s2 values are same, however they have different memory location")
	println(s1 == s2)
	println(s1 == *s3)
	println("Comparing strings using compare method.  Returns 0 if true & -1 if false")
	fmt.Println("Comparing s1 and s2 ", strings.Compare(s1, s2))
	fmt.Println("Comparing s1 & s3 ", strings.Compare(s1, *s3))
	fmt.Println("Comparing s1 with hemnath", strings.Compare(s1, "hemnath"))
	*s3 = "Alia"
	println("After changing value of s3")
	fmt.Println(s1, s2, *s3)

	//compare case-insensitive
	fmt.Println("Case in-sensitive comparision of s1 with hemnath", strings.EqualFold(s1, "hemnath"))

	fmt.Println("===============")
	name := "HemnatH"
	visited := make(map[int32]bool, len(name))
	for i, v := range name {
		fmt.Println(i, v)
		if visited[v] == true {
			fmt.Println("dups")
			fmt.Println(string(v))
		} else {
			visited[v] = true
		}
	}
	fmt.Println("===============")
	for i := range name {
		fmt.Println(string(name[i]))
		kk := name[i]
		fmt.Println(reflect.TypeOf(kk))
	}
	for v1, v2 := range visited {
		println(string(v1), v2)
	}

	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}

}
