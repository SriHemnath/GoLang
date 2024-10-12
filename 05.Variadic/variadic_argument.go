package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println(average(data...))
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
