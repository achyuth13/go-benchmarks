package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workqueue chan Job
	wg        sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workqueue: make(chan Job),
	}
	pool.wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			for job := range pool.workqueue {
				job()
				fmt.Printf("\n")
			}
		}()
	}
	return pool
}

func (p *Pool) AddJob(job Job) {
	p.workqueue <- job
}

func (p *Pool) Wait() {
	close(p.workqueue)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5)

	for i := 0; i < 30; i++ {
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("job : completed %d \n", i)
		}
		pool.AddJob(job)
	}
}
