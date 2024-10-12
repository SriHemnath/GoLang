package main

import "fmt"

func main1() {
	var y int
	for y, z := 1, 1; y < 10; y++ {
		_ = y
		_ = z
	}
	fmt.Println(y) //prints 0

	var m map[string]int //reference type
	println(m)
	//m["a"] = 1 //panic: assignment to entry in nil map
	//fmt.Println(m)

	a := make([]int, 3, 3) //creating slice len=3 & capcity=3
	a[0] = 0
	a[1] = 1
	a[2] = 2
	b := a[:2]                      //copying a reference from 0 to index 2 (ending index is exclusive)
	fmt.Println("a->", a, "b->", b) // "a->", [0,0,0] b-> [0,0]
	b = append(b, 31)               // adding elements to b
	fmt.Println(a, b)               // a=0,0,0  b=[0,0,31]
	b = append(b, 32)               //
	fmt.Println(a, b)               // a=0,0,0  b=[0,0,31,32]
	b = append(b, 33)
	fmt.Println(a, b) // a=0,0,0  b=[0,0,31,32,33]
	//fmt.Println(cap(b))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var arr2 []int //does not allocate memory and points to nil, use append to add values
	arr3 := make([]int, 10)
	elem := copy(arr2, arr) //will not copy as arr2 points to nil
	fmt.Println(arr2, elem)
	elem_2 := copy(arr3, arr[2:5])
	fmt.Println(arr3, elem_2)
	arr2 = append(arr2, arr...) //append function to copy array
	fmt.Println(arr2)
	arr2[0] = 66 //will not have reference of original array
	fmt.Println(arr, arr2)

	aa := make([]int, 1, 1)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(cap(aa))

}

//func f1() { //stack
//	x := S{1}
//	_ = f2(x)
//}
//func f2(x S) *S {
//	y := x //escape to heap
//	return &y
//}
