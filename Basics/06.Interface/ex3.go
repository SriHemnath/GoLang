package main

import "fmt"

type check struct {
}

func main() {
	var i *int        //i=nil
	var j interface{} //j=nil
	var k int         //k=0
	var l *check      //j=nil
	p := new(int)     //p is reference and value inside p is nil
	//*p = 1234
	println("i:")
	do(i)
	println("j:")
	do(j)
	println("k:")
	do(k)
	println("l:")
	do(l)
	println("p:")
	do(p)
	println(i, j, k, l, p)
	fmt.Println(i, j, k, l, p)
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
