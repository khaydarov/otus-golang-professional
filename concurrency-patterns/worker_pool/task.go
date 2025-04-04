package main

import (
	"fmt"
	"time"
)

type Task struct {
	id       int
	fn       func()
	duration time.Duration
}

type TaskResult struct {
	taskId  int
	message string
	err     error
}

func (r TaskResult) String() string {
	return fmt.Sprintf("task %d: %s", r.taskId, r.message)
}
