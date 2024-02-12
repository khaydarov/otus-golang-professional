package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	envDir := args[0]
	envs, err := ReadDir(envDir)
	if err != nil {
		log.Fatalf("Error reading env dir: %v", err)
	}

	cmd := args[1:]

	returnCode := RunCmd(cmd, envs)

	os.Exit(returnCode)
}
