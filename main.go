package main

import (
	"fmt"
	"os/exec"
)

// a cli program to read file with a given name

func getGitVersion() string {
	// excute "git describe --tag --always" in terminal

	cmd := exec.Command("pwd")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return "unknown"
	}

	return string(out)
}

func main() {
	// app := cli.NewApp()
	// app.Name = "golang-cli-example"
	// app.Usage = "A simple cli program to read file with a given name"
	// app.Version = "1.0.0"
	// app.Action = func(c *cli.Context) error {
	// 	// do something
	// 	return nil
	// }
	// app.Run([]string{"golang-cli-example", "--help"})

	fmt.Println(getGitVersion())
}
