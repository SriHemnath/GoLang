package main

import (
	"fmt"
	"math"
	"sync"
)

//print list armstrong numbers 0 to 1 lakh

func main() {
	printArmstrongNumbers()
}

func printArmstrongNumbers() {

	result := make(chan int)
	quit := make(chan int)
	var wg sync.WaitGroup
	go printResult(result, quit)
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go checkArmstrong(i, result, &wg)
	}
	close(result)
	fmt.Println("Channel closed, ", <-quit)
	wg.Wait()
}

func printResult(result chan int, quit chan int) {
	for {
		select {
		case num, ok := <-result:
			if !ok {
				quit <- 0
			}
			fmt.Println(num)
		}
	}
}

func checkArmstrong(num int, finalResult chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	length := getNumLength(num)
	tempNum := num
	var result int
	for i := 0; i < length; i++ {
		reminder := tempNum % 10
		result += int(math.Pow(float64(reminder), float64(length)))
		tempNum /= 10
	}
	if num == result {
		finalResult <- num
	}
}

func getNumLength(num int) (count int) {
	for num != 0 {
		count++
		num /= 10
	}
	return count
}
