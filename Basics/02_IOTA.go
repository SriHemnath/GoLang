package main

import "fmt"

const(
	a = iota +1
	b
	c
)

const(
	d,e,f=iota,iota,iota
)

func main(){
	fmt.Println(a,b,c,d,e,f)
	fmt.Println("IOTA starts from 0 and increment value after each specification. Values are reset in new declaration")

	//only IF & FOR loop available in GO.
	//using for as while
	v:=1
	for  v < 10  {
		fmt.Println("something")
		v++
	}

	//we range over a collection of values using for range
	s:=[]int{1,2,3,4,5,6}
	for _, val := range s {
		fmt.Println(val)
	}

	fmt.Println("Printing prime numbers")
	printPrime()
}

//PRINT PRIME
func printPrime(){
	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) bool{
	if n<=1{
		return false
	}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}			
	}
	return true
}