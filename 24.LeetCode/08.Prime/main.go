package main

import (
	"fmt"
	"sync"
)

func main() {
	primeOrNot(30)
}

func primeOrNot(num int) {
	var wg sync.WaitGroup
	result := make([]int, 0)
	isPrime := make(chan int)
	go getPrimeValues(isPrime, &result)
	for i := 2; i < num; i++ {
		wg.Add(1)
		go checkPrime(i, isPrime, &wg)
	}

	wg.Wait()
	close(isPrime)
	fmt.Println("List of prime numbers:", result)
}

func getPrimeValues(isPrime chan int, result *[]int) {
	for i := range isPrime {
		*result = append(*result, i)
	}
}

func checkPrime(num int, prime chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("is prime, ", num)
		prime <- num
	}
}
