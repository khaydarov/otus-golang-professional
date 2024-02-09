package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	c := exec.Command(cmd[0], cmd[1:]...)

	c.Env = os.Environ()
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
			continue
		}
		c.Env = append(c.Env, k+"="+v.Value)
	}

	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		return 1
	}

	return 0
}
