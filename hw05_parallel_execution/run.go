package hw05parallelexecution

import (
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

	// run goroutines
	var countOfErrors int32
	for workerID := 0; workerID < n; workerID++ {
		go runTask(workerID, wg, tasksChannel, &countOfErrors)
	}

	// send tasks to workers
	for _, task := range tasks {
		tasksChannel <- task

		if atomic.LoadInt32(&countOfErrors) >= maxErrorCount {
			// if count of errors exceeds limit then close channel and stop all workers
			close(tasksChannel)
			wg.Wait()
			return ErrErrorsLimitExceeded
		}
	}

	close(tasksChannel)
	wg.Wait()
	return nil
}

// runTask consumes task and runs it
// every runner shares `countOfErrors` variable to check if errors exceeds max errors count.
func runTask(
	workerID int,
	wg *sync.WaitGroup,
	tasksChannel <-chan Task,
	countOfErrors *int32,
) {
	defer wg.Done()

	log.Printf("Worker %d started\n", workerID)
	for task := range tasksChannel {
		err := task()
		if err != nil {
			atomic.AddInt32(countOfErrors, 1)
		}

		log.Printf("Worker %d completed task\n", workerID)
	}
}
