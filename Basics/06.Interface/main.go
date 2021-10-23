package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

type Animal interface {
	speak() string
}

type Square struct {
	side float64
}

type Rectangle struct {
	side float64
}

func info(z Shape) {
	fmt.Println(z.area())
}

//Square is an shape now since the method signature matches the shape interface
//now you can pass square values to shape type since square is a shape
func (s Square) area() float64 {
	return s.side * s.side
}

func area() float64 {
	return 2.0 * 2.0
}

func (r Rectangle) area() float64 {
	return r.side * 2
}

type Indians struct {
}

type Americans struct{}

type Chinese struct{}

func (i Indians) speak() string {
	return "Vanthae Matharam"
}

func (a Americans) speak() string {
	return "Uptown Funk"
}

func (c Chinese) speak() string {
	return "eat everything"
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {

	s := Square{10}
	info(s)
	r := Rectangle{4}
	fmt.Println(r.area())
	info(r)

	animals := []Animal{Indians{}, Americans{}, Chinese{}}
	for _, v := range animals {
		fmt.Println(v.speak())
	}

	names := []string{"stanley", "david", "oscar"}
	//PrintAll(names)   will not work
	vals := make([]interface{}, len(names))
	fmt.Println("Type of vals ", fmt.Sprintf("%T", vals))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}
