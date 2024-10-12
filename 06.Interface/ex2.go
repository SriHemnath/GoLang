package main

import "fmt"

func printAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func printThis(vals interface{}) {
	fmt.Println(vals)
	kk := vals.([]string)
	fmt.Println(kk)
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	//printAll(names)  //this will throw error - Cannot use 'names' (type []string) as the type []interface{}
	printThis(names)
}
