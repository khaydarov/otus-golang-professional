package main

import (
	"bytes"
	"errors"
	"io"
	"log"
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
			log.Printf("Invalid file name: %s", entry.Name())

			continue
		}

		v, err := readValue(dir + "/" + entry.Name())
		needToRemove := false
		if err != nil {
			needToRemove = true
		}
		envs[entry.Name()] = EnvValue{Value: v, NeedRemove: needToRemove}
	}

	return envs, nil
}

func readValue(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	data, _ := io.ReadAll(file)
	if len(data) == 0 {
		return "", errors.New("empty file")
	}
	buf := bytes.NewBuffer(data)
	data, _ = buf.ReadBytes('\n')
	data = bytes.ReplaceAll(data, []byte("\x00"), []byte("\n"))

	return strings.TrimRight(string(data), " \t\r\n"), nil
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
