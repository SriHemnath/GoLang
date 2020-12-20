package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("Number of go routines = ",runtime.NumGoroutine())
	ca := make(chan int)
	go func(){
		ca <- 786
		ca <- 007
		ca <- 2
	}()
	
	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("Number of go routines = ",runtime.NumGoroutine())

	ca1 := make(chan string, 1)
	ca1 <- "Hemnath"
	fmt.Println(<-ca1)

}