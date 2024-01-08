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

	tasksChannel := make(chan Task, len(tasks))

	// calc worker count: if len(tasks) are less than N then we can run len(tasks) goroutines
	workerCount := n
	if workerCount > len(tasks) {
		workerCount = len(tasks)
	}

	wg := &sync.WaitGroup{}
	wg.Add(workerCount)

	// run goroutines
	var countOfErrors int32
	for workerID := 0; workerID < workerCount; workerID++ {
		go runTask(workerID, wg, tasksChannel, &countOfErrors, maxErrorCount)
	}

	// send tasks to workers
	for _, task := range tasks {
		tasksChannel <- task
	}
	close(tasksChannel)

	wg.Wait()
	if atomic.LoadInt32(&countOfErrors) >= maxErrorCount {
		return ErrErrorsLimitExceeded
	}

	return nil
}

// runTask consumes task and runs it
// every runner shares `countOfErrors` variable to check if errors exceeds max errors count.
func runTask(workerID int, wg *sync.WaitGroup, tasksChannel <-chan Task, countOfErrors *int32, maxErrorCount int32) {
	defer wg.Done()

	log.Printf("Worker %d started\n", workerID)
	for task := range tasksChannel {
		if atomic.LoadInt32(countOfErrors) >= maxErrorCount {
			log.Printf("Errors count exceeded. Worker %d terminating\n", workerID)
			return
		}

		err := task()
		if err != nil {
			atomic.AddInt32(countOfErrors, 1)
		}

		log.Printf("Worker %d completed task\n", workerID)
	}
}
