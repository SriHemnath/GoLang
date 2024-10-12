package main

import (
	"GoLang/Basics/01_Scope/vis"
	"fmt"
)

//a has package level scope
var a = 0

//anymodifications done to a will reflect in other places too
//declare it as const to avoid modifications in different places
func increment() int {
	a++
	return a
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(a)

	//go mod init and then import the package to use the func & variables of it
	fmt.Println(vis.Name, vis.Location)
	vis.Printer(5)

	emp := vis.NewEmployee()
	fmt.Println(emp)
	salary := emp.GetSalary()
	fmt.Println(salary)

}
