package vis

import "fmt"

var Name = "Hemnath"

const Location = "chennai"

// Employee value object: immutable struct.  Does not have any unique ID
type Employee struct {
	name   string
	id     int
	salary float64
}

func Printer(i int) {
	fmt.Println(i)
}

func NewEmployee() *Employee {
	return &Employee{
		name:   "Hemnath",
		id:     1,
		salary: 50000,
	}
}

func (emp *Employee) GetSalary() float64 {
	return emp.salary
}
