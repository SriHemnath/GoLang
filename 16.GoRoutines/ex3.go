package main

import (
	"fmt"
	"sync"
)

var N int = 10

/*Go routines run concurrently and for loop completes before the go routines starts execution. Running it every time produces different result
We can use time.Sleep after go routine so it gets value from for loop or pass the value from loop as parameter to routines.*/
func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(len(m))
	fmt.Println(m)
	fixFuncLiteral()
	syncMap()
}

//fix 1
func fixFuncLiteral() {
	println("Passing value in function literal")
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(len(m))
	fmt.Println(m)
}

//fix 2
/*Use a temporary variable to and use that variable*/
func syncMap() {
	println("Using sync map")
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	var contexts = sync.Map{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		j := i
		go func() {
			mu.Lock()
			contexts.Store(j, j)
			mu.Unlock()
			defer wg.Done()
		}()
	}

	wg.Wait()
	contexts.Range(func(k, v interface{}) bool {
		fmt.Println("range (): ", k, v)
		return true
	})
}
