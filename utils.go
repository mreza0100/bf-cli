package main

import (
	"os"
	"os/exec"
)

func isInArgs(args []string, values []string) (isIn bool, next []string) {
	for idx, a := range args {
		for _, v := range values {
			if a == v {
				return true, args[idx+1:]
			}
		}
	}

	return false, args
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
