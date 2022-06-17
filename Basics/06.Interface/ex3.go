package main

import "fmt"

type check struct {
}

func main() {
	var i *int        //i=nil
	var k int         //k=0
	var j interface{} //j=nil
	var l *check      //j=nil
	p := new(int)     //p is reference and value inside p is nil
	//*p = 1234
	do(i)
	do(j)
	do(k)
	do(l)
	do(p)
	println(i, j, k, l, p)
	fmt.Println(i, j, k, l, p)
}

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
