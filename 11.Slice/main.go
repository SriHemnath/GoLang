package main

import (
	"fmt"
	"reflect"
)

/*
append works different with different data types.
For ex. string is always doubled when exceeded capacity.
[]int capacity is increased based on the number of elements added
*/
func main() {

	//CREATING SLICE
	//1. DECLARE EMPTY SLICE
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice))
	fmt.Println(reflect.ValueOf(strSlice).Type()) //print []string
	fmt.Println(reflect.ValueOf(strSlice).Kind()) //print slice

	fmt.Println(intSlice) //print empty slice
	intSlice = append(intSlice, 12)
	fmt.Println(intSlice, len(intSlice), cap(intSlice))

	//2.UING MAKE KEYWORD
	var iSlice = make([]int, 0, 3)     // when length and capacity is same
	var sSlice = make([]string, 0, 20) // when length and capacity is different
	iSlice = append(iSlice, 1, 2, 3)
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(iSlice), cap(iSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(sSlice), cap(sSlice))
	fmt.Println(sSlice)

	iSlice = append(iSlice, 4, 5, 6, 7, 8, 9)
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(iSlice), cap(iSlice))

	//3.UING SLICE LITERAL
	//var intSlice = []int{10, 20, 30, 40}

	//4.USING NEW KEYWORD
	var nSlice = new([10]int)[0:5] //length = 5 capacity = 10
	fmt.Printf("nSlice \tLen: %v \tCap: %v\n", len(nSlice), cap(nSlice))
	nSlice = append(nSlice, 6, 7, 8, 9, 10, 11)
	fmt.Println(nSlice)
	fmt.Printf("nSlice \tLen: %v \tCap: %v\n", len(nSlice), cap(nSlice))
	nSlice = append(nSlice, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Printf("nSlice \tLen: %v \tCap: %v\n", len(nSlice), cap(nSlice))

	var someValue interface{} = 2
	someValue = 2.0
	someValue = "Hemnath"
	fmt.Println(someValue)
	fmt.Println(reflect.TypeOf(someValue))

	//reverse the order of slice
	x := []int{5, 3, 1, 4, 2}
	fmt.Println("Original array: ", x)
	reverse(x)
	fmt.Println("After sorting: ", x)
	modify(x)
	fmt.Println("After modifying: ", x) //this will modify original array
}

func reverse(x []int) {
	for a, b := 0, len(x)-1; a < b; a, b = a+1, b-1 {
		x[a], x[b] = x[b], x[a]
	}
	x = append(x, 99) // won't modify original array
	// x = []int{9, 7, 6, 5, 4}   //will not modify actual x
}

func modify(x []int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
}
