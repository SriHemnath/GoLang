package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	quit := make(chan struct{})

	var wg sync.WaitGroup
	//wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go doSomething(&wg, quit, ctx)
	go doAnotherthing(&wg, quit, ctx)
	time.Sleep(5 * time.Second)
	//wg.Wait()
	//quit <- struct{}{}
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("completed")
}

func doSomething(wg *sync.WaitGroup, quit chan struct{}, ctx context.Context) {
	for {
		select {
		case <-quit:
			fmt.Println("received quit in something")
			//wg.Done()
			return
		case <-ctx.Done():
			fmt.Println("received done context in something")
			return
		}
	}

}

func doAnotherthing(wg *sync.WaitGroup, quit chan struct{}, ctx context.Context) {
	for {
		select {
		case <-quit:
			fmt.Println("received quit in another thing")
			//wg.Done()
			return
		case <-ctx.Done():
			fmt.Println("received done context in another thing")
			return

		}
	}
}
