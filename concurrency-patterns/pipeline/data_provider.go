package main

import (
	"errors"
	"math/rand"
)

var ErrNoMoreInts = errors.New("no more ints")

func createIntStream(terminate <-chan struct{}, fn func() (int, error)) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-terminate:
				return
			default:
				v, err := fn()
				if errors.Is(err, ErrNoMoreInts) {
					return
				}

				out <- v
			}
		}
	}()

	return out
}

func RandomInt() (int, error) {
	return rand.Intn(1000), nil
}

func DeterministicInts(ints []int) func() (int, error) {
	cursor := 0
	return func() (int, error) {
		defer func() { cursor++ }()
		if cursor == len(ints) {
			return 0, ErrNoMoreInts
		}
		return ints[cursor], nil
	}
}

func NewRandomIntStream(terminate <-chan struct{}) Out {
	return createIntStream(terminate, RandomInt)
}

func NewDeterministicIntStream(terminate <-chan struct{}, ints []int) Out {
	return createIntStream(terminate, DeterministicInts(ints))
}
