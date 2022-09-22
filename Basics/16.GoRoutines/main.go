package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	quit := make(chan bool) //no buffer size-blocking channel.  Will wait till receiver is ready
	fmt.Println("Hemnath")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			fmt.Println("Outside select")
			select {
			case <-quit:
				return
			default:
				fmt.Println("Inside default.  Waiting...")
			}
		}
	}()
	// …
	time.Sleep(5 * time.Second)
	quit <- true //commenting this will put go routine in infinite waiting loop
	wg.Wait()
	close(quit)
	next := incrementor() // next is a function returned by incrementor
	next1 := incrementor()
	fmt.Println(next())  // prints 1
	fmt.Println(next())  // prints 2
	fmt.Println(next())  // prints 3
	fmt.Println(next1()) // prints 1
	fmt.Println(next1()) //print 2
	fmt.Println(next())  //print 4
}

//function closure
func incrementor() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
