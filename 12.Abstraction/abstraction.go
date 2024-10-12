package main

import "fmt"

//hiding the implementation details and showing only functionality to the user
//abstraction can only be implemented using methods(func with receiver)

//Metho sets: A non-pointer receiver accepts both non-pointer and pointer value, however any modifications
//            done within func will not reflect outside cause a copy of value is passed
//            A pointer receiver works only with pointer values & changes inside method with reflect everywhere
type Laptop interface {
	Spec() string
}

type Lenovo struct {
	model     string
	name      string
	processor string
}

type Apple struct {
	model     string
	name      string
	processor string
}

func (l *Lenovo) Spec() string {
	return fmt.Sprintf("Name:%s, Model:%s, Processor:%s ", l.name, l.model, l.processor)
}

func (a Apple) Spec() string {
	return fmt.Sprintf("Name:%s, Model:%s, Processor:%s ", a.name, a.model, a.processor)
}

func main() {
	var i Laptop
	l := &Lenovo{
		model:     "GT",
		processor: "RTX3080",
		name:      "Legion",
	}
	i = l
	fmt.Println(l.Spec())
	fmt.Println(i.Spec())

	j := Lenovo{
		model:     "MCLaren",
		processor: "RTX3070",
		name:      "Legion",
	}
	fmt.Println("Non-pointer ", j.Spec()) // go takes care of converting it to pointer
	// i = j won't work, cause Spec method is attached with pointer receiver. Compiler can't find method with non-pointer receiver
	// i = &j will work

	a := Apple{
		model:     "GT",
		processor: "M1",
		name:      "MacbookAir",
	}

	i = &a
	fmt.Println(i.Spec())
	fmt.Println(a.Spec())
}
