package main

import (
	"os"
	"os/exec"
)

func getGitVersion() string {
	cmd := exec.Command("git", "describe", "--tag", "--always")

	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(stdout)
}

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
