package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 5)
	wg.Add(1)
	for i := 1; i < 5; i++ {
		ch <- i
	}
	go access(ch)
	close(ch)
	wg.Wait()
}

func access(ch chan int) {
	defer wg.Done()
	for {
		if val, open := <-ch; open {
			fmt.Println(val)
		} else {
			fmt.Println("channel closed")
			break
		}
	}
}
