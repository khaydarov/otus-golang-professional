package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envFiles := []struct {
		name    string
		content string
	}{
		{
			name:    "APP_PORT",
			content: "8080",
		},
		{
			name:    "APP_HOST",
			content: "localhost",
		},
		{
			name:    "APP_DEBUG",
			content: "true",
		},
	}

	_ = os.Mkdir("./testdata/temp/", os.ModePerm)
	for _, envFile := range envFiles {
		file, _ := os.Create("./testdata/temp/" + envFile.name)
		file.WriteString(envFile.content)
	}

	envs, err := ReadDir("./testdata/temp/")
	require.NoError(t, err)
	require.Equal(t, envs["APP_PORT"].Value, "8080")
	require.Equal(t, envs["APP_HOST"].Value, "localhost")
	require.Equal(t, envs["APP_DEBUG"].Value, "true")

	os.RemoveAll("./testdata/temp/")
}
