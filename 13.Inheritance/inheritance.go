package main

import "fmt"

// Object acquiring the properties & behavior of another object is inheritance.
// GO does not have object & hence doesn't support inheritance, however achived through
// struct using composition (struct having another struct)

// In composition, base structs can be embedded into a child struct and the methods of the
// base struct can be directly called on the child struct

type BaseOne struct {
	base_one string
}

type BaseTwo struct {
	base_two string
}

type Child struct {
	BaseOne
	BaseTwo
}

func (b *BaseOne) print() {
	fmt.Println(b.base_one)
}

func (b *BaseTwo) print() {
	fmt.Println(b.base_two)
}

func (b *BaseOne) printBaseOne() {
	fmt.Println(b.base_one)
}

func main() {

	c := Child{
		BaseOne{
			base_one: "This is from base one",
		},
		BaseTwo{
			base_two: "This is from base two",
		},
	}
	c.printBaseOne()  //Child has access to parent methods
	c.BaseTwo.print() //Since method names are same, calling using dot operator
	c.BaseOne.print()
}
