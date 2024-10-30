package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Hemnath", "you passed value")
	ctx = context.WithValue(ctx, "Samantha", "I am awesome")
	doSomething(ctx)
	fmt.Println(ctx.Value("Hemnath"))

}

func doSomething(ctx context.Context) {
	fmt.Println("something")
	fmt.Println(ctx.Value("Hemnath"))
	fmt.Println(ctx.Value("Samantha"))
	ctx = context.WithValue(ctx, "Pallavi", "I am legend") //new context can be passed to children, but won't reflect in parent
	doOneMore(ctx)
	nums := make(chan int)
	ctx, cancelContext := context.WithCancel(ctx)
	go printNums(ctx, nums)
	for i := 0; i <= 3; i++ {
		nums <- i
	}
	cancelContext()
	time.Sleep(100 * time.Millisecond)
}

func doOneMore(ctx context.Context) {
	fmt.Println("from doonemore")
	fmt.Println(ctx.Value("Pallavi"))
	fmt.Println(ctx.Value("Samantha"))
}

func printNums(ctx context.Context, nums <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("done")
			return
		case val := <-nums:
			fmt.Println(val)
		}
	}
}
