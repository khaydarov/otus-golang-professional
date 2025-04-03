package main

import (
	"context"
)

type PromiseResult struct {
	Value interface{}
	Err   error
}

func Promise(ctx context.Context, task func() (interface{}, error)) <-chan PromiseResult {
	ch := make(chan PromiseResult, 1)

	go func() {
		defer close(ch)

		if ctx.Err() != nil {
			ch <- PromiseResult{Err: ctx.Err()}
			return
		}

		taskResult := make(chan PromiseResult, 1)
		go func() {
			defer close(taskResult)

			value, err := task()
			taskResult <- PromiseResult{Value: value, Err: err}
		}()

		select {
		case <-ctx.Done():
			ch <- PromiseResult{Err: ctx.Err()}
		case result := <-taskResult:
			ch <- result
		}
	}()

	return ch
}
