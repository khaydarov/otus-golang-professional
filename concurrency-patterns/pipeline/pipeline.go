package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/khaydarov/otus-golang-professional/concurrency-patterns/pipeline/transformers"
)

type Pipeline struct {
	terminate <-chan struct{}
}

func NewPipeline(terminate <-chan struct{}) Pipeline {
	return Pipeline{terminate: terminate}
}

func (p Pipeline) Execute(in In, fns ...transformers.Transformer) Out {
	result := in
	for _, transformer := range fns {
		name := runtime.FuncForPC(reflect.ValueOf(transformer).Pointer()).Name()
		result = newStage(p.terminate, name, transformer)(result)
	}
	return result
}

func main() {
	terminate := make(chan struct{})
	defer close(terminate)
	intStream := NewDeterministicIntStream(terminate, []int{1, 2, 3, 4, 5})

	pipeline := NewPipeline(terminate)
	result := pipeline.Execute(
		intStream,
		transformers.Incrementor,
		transformers.Multiplier(3),
	)

	for v := range result {
		fmt.Println("transformed value", v)
	}
}
