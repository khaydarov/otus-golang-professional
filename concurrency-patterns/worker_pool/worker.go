package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id             int
	idle           bool
	tasksProcessed int
	totalWorkTime  time.Duration
	lastTaskTime   time.Time
}

func (w *Worker) GetMetrics() string {
	return fmt.Sprintf("Worker %d: %d tasks processed, %s total work time", w.id, w.tasksProcessed, w.totalWorkTime)
}

func (w *Worker) Handle(
	ctx context.Context,
	tasks <-chan Task,
	result chan<- TaskResult,
	wg *sync.WaitGroup,
) {
	w.idle = false
	defer func() {
		w.idle = true
		wg.Done()

		fmt.Printf("%s\n", w.GetMetrics())
	}()

	for {
		select {
		case <-ctx.Done():
			result <- TaskResult{
				taskId:  0,
				message: fmt.Sprintf("Worker %d stopped: %v", w.id, ctx.Err()),
				err:     ctx.Err(),
			}
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				result <- TaskResult{
					taskId:  task.id,
					message: fmt.Sprintf("Task %d cancelled", task.id),
					err:     ctx.Err(),
				}
				return
			case <-time.After(task.duration):
				result <- TaskResult{
					taskId:  task.id,
					message: fmt.Sprintf("Completed by worker %d", w.id),
					err:     nil,
				}
			}

			w.tasksProcessed++
			w.totalWorkTime += time.Since(w.lastTaskTime)
			w.lastTaskTime = time.Now()
		}
	}
}
