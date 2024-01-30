package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	err := Copy("./testdata/input.txt", "./testdata/out.txt", 7000, 0)
	require.True(t, errors.Is(err, ErrOffsetExceedsFileSize), "must return offset error")
}
