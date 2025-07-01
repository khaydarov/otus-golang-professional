package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	returnCode := RunCmd([]string{"echo", "hello"}, Environment{})
	require.Equal(t, 0, returnCode)
}
