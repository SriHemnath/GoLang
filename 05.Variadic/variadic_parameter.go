package main

import "fmt"

func main() {
	n := avaerage(13, 3, 4, 5, 6)
	fmt.Println(n)
	fmt.Println(avaerage(5))
}

func avaerage(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
