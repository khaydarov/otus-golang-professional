package main

import (
	"github.com/khaydarov/otus-golang-professional/concurrency-patterns/pipeline/transformers"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) Out

func newStage(terminate <-chan struct{}, name string, fn transformers.Transformer) Stage {
	return func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)

			for v := range in {
				select {
				case <-terminate:
					return
				case out <- fn(v):
				}
			}
		}()

		return out
	}
}
