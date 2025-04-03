package main

import "context"

func Generator(ctx context.Context, data []interface{}) <-chan interface{} {
	ch := make(chan interface{})

	if ctx.Err() != nil {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)

		for _, v := range data {
			if ctx.Err() != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}
