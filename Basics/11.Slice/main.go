package main

import (
	"fmt"
	"reflect"
)

func main() {

	//CREATING SLICE
	//1. DECLARE EMPTY SLICE
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Type())
	fmt.Println(reflect.ValueOf(strSlice).Type())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	//2.UING MAKE KEYWORD
	var iSlice = make([]int, 10)        // when length and capacity is same
	var sSlice = make([]string, 10, 20) // when length and capacity is different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(iSlice), cap(iSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(sSlice), cap(sSlice))

	//3.UING SLICE LITERAL
	//var intSlice = []int{10, 20, 30, 40}

	//4.USING NEW KEYWORD
	var nSlice = new([10]int)[0:5] //length = 10 capacity = 50
	fmt.Printf("nSlice \tLen: %v \tCap: %v\n", len(nSlice), cap(nSlice))
	nSlice = append(nSlice, 11, 12, 13, 14, 15, 16)
	fmt.Println(nSlice)
	fmt.Printf("nSlice \tLen: %v \tCap: %v\n", len(nSlice), cap(nSlice))

	var someValue interface{} = 2
	someValue = 2.0
	someValue = "Hemnath"
	fmt.Println(someValue)
	fmt.Println(reflect.TypeOf(someValue))
}
