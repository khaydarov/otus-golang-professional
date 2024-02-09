package main

import (
	"bytes"
	"fmt"
	"io"
	log "log/slog"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)
	for _, entry := range entries {
		if containsInvalidSymbol(entry.Name()) {
			log.Info(fmt.Sprintf("Invalid file name: %s", entry.Name()))

			continue
		}

		v := readValue(dir + "/" + entry.Name())
		needToRemove := false
		if len(v) == 0 {
			needToRemove = true
		}
		envs[entry.Name()] = EnvValue{Value: v, NeedRemove: needToRemove}
	}

	return envs, nil
}

func readValue(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}

	defer file.Close()

	data, _ := io.ReadAll(file)
	buf := bytes.NewBuffer(data)
	data, _ = buf.ReadBytes('\n')
	data = bytes.ReplaceAll(data, []byte("\x00"), []byte("\n"))

	line := strings.TrimRight(string(data), "\n")
	return strings.TrimRight(line, " ")
}

func containsInvalidSymbol(s string) bool {
	for _, v := range s {
		if isInvalidSymbol(v) {
			return true
		}
	}

	return false
}

func isInvalidSymbol(b rune) bool {
	return b == '='
}
