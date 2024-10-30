package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id      int
	jobchan chan int
}

type Pool struct {
	workers  []Worker
	jobQueue chan int
}

func (w Worker) processWork() {
	for job := range w.jobchan {
		fmt.Printf("FROM WORKER %d value %d\n", w.id, job) //common framework
		time.Sleep(time.Second * 1)
	}
}

func NewPool(numWorkers int) *Pool {
	jobQueue := make(chan int)
	workers := make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker := Worker{id: i, jobchan: make(chan int)}
		workers[i] = worker
		go worker.processWork()
	}
	return &Pool{workers: workers, jobQueue: jobQueue}
}

func (p *Pool) Start() {
	for _, worker := range p.workers {
		go func(worker Worker) {
			fmt.Println("I am started:", worker.id)
			for job := range p.jobQueue {
				// fmt.Println("running")
				worker.jobchan <- job
			}
			fmt.Println("I am finished:", worker.id)
		}(worker)
	}
}

func main() {
	//1. Simple approach
	workerPool := NewPool(2)
	workerPool.Start()

	for i := 0; i < 10; i++ {
		fmt.Println("sending to channel:", i)
		workerPool.jobQueue <- i
		fmt.Println("sent to channel:", i)
	}
	time.Sleep(time.Second * 10)

	close(workerPool.jobQueue)
	time.Sleep(2)

	//2. using errorgroup from goang/x/..
}
