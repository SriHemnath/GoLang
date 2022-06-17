package main

type Person struct {
	Name string
	Age  int
}

//changes won't be reflected, will have a copy of values and
func (p Person) modifyAge() {
	p.Age = 22
}

//changes will be reflected
func (p *Person) modifyName() {
	p.Name = "Biju"
}

func main() {
	p := Person{Name: "Hemnath", Age: 27}
	p.modifyAge()

	k := Person{Name: "Samantha", Age: 22}
	k.modifyName()
	println(p.Name, p.Age)
	println(k.Name, k.Age)
}
