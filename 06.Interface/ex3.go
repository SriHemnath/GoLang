package main

import (
	"fmt"
)

type check struct {
}

func main() {
	var i *int        //i=nil
	var j interface{} //j=nil
	var k int         //k=0
	var l *check      //j=nil
	var ll check
	var m = l
	p := new(int) //p is reference and value inside p is nil
	//*p = 1234
	println("i: ", i, i == nil)
	do(i)
	println("j:")
	do(j)
	println("k:")
	do(k)
	println("l:")
	do(l)
	println("ll:")
	do(ll)
	println("p:")
	do(p)
	println("m:")
	do(m)
	println(i, j, k, l, p, m)
	fmt.Println(i, j, k, l, ll, p, m)
	// fmt.Println(unsafe.Sizeof(ll))
	// fmt.Println(unsafe.Sizeof(l))
	// fmt.Println(unsafe.Sizeof(j))
}

//whenever interface receives a value, it will have reference of that value as interface{} is a reference type.
//if empty interface is passed then it will be nil
func do(i interface{}) {
	//var pp *interface{}
	//println(pp)
	println("Inside do ", i)
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
