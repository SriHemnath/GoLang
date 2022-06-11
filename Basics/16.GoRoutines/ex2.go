package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 5)
	wg.Add(1)
	go access(ch)

	for i := 0; i < 9; i++ {
		fmt.Println("waiting")
		ch <- i
		fmt.Println("Filled")
	}
	close(ch)
	wg.Wait()
	fmt.Println("finished main")
}

func access(ch chan int) {
	defer wg.Done()
	fmt.Println("start accessing channel\n")
	//for i := range ch {
	//	fmt.Println(i)
	//}
	for {
		if a, b := <-ch; b {
			fmt.Println(b)
			fmt.Println(a)
		} else {
			fmt.Println(b)
			break
		}
	}
}
