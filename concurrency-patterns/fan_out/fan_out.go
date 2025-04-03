package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	ID       int
	Duration time.Duration
}

func TaskGenerator() chan Task {
	ch := make(chan Task)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			duration := time.Duration(rand.Intn(1000)) * time.Millisecond
			ch <- Task{ID: i, Duration: duration}
		}
	}()
	return ch
}

func taskExecutor(executorID int, tasks <-chan Task) <-chan struct{} {
	result := make(chan struct{})
	go func() {
		defer close(result)
		for task := range tasks {
			fmt.Printf("Executor %d started task %d, duration: %s\n", executorID, task.ID, task.Duration)
			time.Sleep(task.Duration)
			fmt.Printf("Executor %d finished task %d\n", executorID, task.ID)
		}
	}()
	return result
}

func fanOut(workers int, tasks <-chan Task) []<-chan struct{} {
	result := make([]<-chan struct{}, workers)

	for i := 0; i < workers; i++ {
		result[i] = taskExecutor(i, tasks)
	}

	return result
}

func main() {
	tasks := TaskGenerator()
	results := fanOut(3, tasks)

	for _, result := range results {
		<-result
	}
}
