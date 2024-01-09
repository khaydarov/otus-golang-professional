package hw05parallelexecution

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	maxErrorCount := int32(m)
	if maxErrorCount <= 0 {
		return ErrErrorsLimitExceeded
	}

	tasksChannel := make(chan Task)
	wg := &sync.WaitGroup{}
	wg.Add(n)

	ctx, cancel := context.WithCancel(context.Background())

	// run goroutines
	var countOfErrors int32
	for workerID := 0; workerID < n; workerID++ {
		go runTask(ctx, workerID, wg, tasksChannel, &countOfErrors, maxErrorCount)
	}

	// send tasks to workers
	finishWithError := false
	for _, task := range tasks {
		tasksChannel <- task

		if countOfErrors >= maxErrorCount {
			finishWithError = true
			break
		}
	}
	close(tasksChannel)

	cancel()
	wg.Wait()

	if finishWithError {
		return ErrErrorsLimitExceeded
	}

	return nil
}

// runTask consumes task and runs it
// every runner shares `countOfErrors` variable to check if errors exceeds max errors count.
func runTask(
	ctx context.Context,
	workerID int,
	wg *sync.WaitGroup,
	tasksChannel <-chan Task,
	countOfErrors *int32,
	maxErrorCount int32,
) {
	defer wg.Done()

	log.Printf("Worker %d started\n", workerID)
	for {
		select {
		case task := <-tasksChannel:
			if atomic.LoadInt32(countOfErrors) >= maxErrorCount {
				log.Printf("Errors count exceeded. Worker %d terminating\n", workerID)
				return
			}

			if task == nil {
				continue
			}

			err := task()
			if err != nil {
				atomic.AddInt32(countOfErrors, 1)
			}

			log.Printf("Worker %d completed task\n", workerID)
		case <-ctx.Done():
			log.Printf("Stop signal. Worker %d terminating\n", workerID)
			return
		}
	}
}
