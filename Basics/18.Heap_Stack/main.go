package main

import (
	"fmt"
)

type hem struct {
	name string
	age  int
}

/*Use the following commands to find the "escape sequence" of values.
If the compiler is not able to prove the value is only used in func it will push to heap memory
1. If value is large
2. If pointer value from func is referenced outside(pointer value created and returned). Referred as "Pointer escape"
3. Dynamic Type escape - If parameter type is interface, difficult to determine type at compile-time it escapes to heap. Ex.fmt.Println
*/

//go build -gcflags '-m' ./ex3.go
//go tool compile -m ex3.go
func main() {
	var i = 0
	//fmt.Println(i)    //escapes to heap
	i++
	print(i)
	//log.Print(i)		//escapes to heap
	m := make(map[string]bool) //escapes to heap
	mn := new(map[string]bool) //escapes to heap
	var v map[string]bool
	fmt.Sprintf("%T%T%T", m, mn, v)
	println(m, mn, v)
	sdf := new(hem)
	//fmt.Println(*sdf)
	qe := new(int)
	println(*qe)
	sdf.age = 14
	println(sdf)
	asdf := hem{}
	asdf.age = 14
	//fmt.Println(asdf)
}
