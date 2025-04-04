package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type WorkerPool struct {
	count   int
	workers []*Worker

	wg sync.WaitGroup
}

func NewWorkerPool(count int) *WorkerPool {
	pool := &WorkerPool{
		count:   count,
		workers: make([]*Worker, count),
	}
	for i := 0; i < pool.count; i++ {
		pool.workers[i] = &Worker{id: i, idle: true, lastTaskTime: time.Now()}
	}

	return pool
}

func (p *WorkerPool) Handle(stop <-chan struct{}, tasks <-chan Task) <-chan TaskResult {
	ctx, cancel := context.WithCancelCause(context.Background())
	results := make(chan TaskResult, p.count*2)

	go func() {
		<-stop
		cancel(errors.New("stop signal received"))
	}()

	go func() {
		defer func() {
			close(results)
		}()

		p.wg.Add(p.count)
		for i := 0; i < p.count; i++ {
			worker := p.workers[i]
			go worker.Handle(ctx, tasks, results, &p.wg)
		}

		p.wg.Wait()
	}()

	return results
}
