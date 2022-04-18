package main

import (
	"os"

	"github.com/urfave/cli"
)

// a cli program to read file with a given name

var app = cli.NewApp()

func main() {
	app.Name = "brainfuck-cli"
	app.Usage = "brainfuck-cli"
	app.Version = getGitVersion()
	app.Author = "M.Reza khosravi"

	actions := new(Actions)
	app.Commands = []cli.Command{
		{
			Name:    "file",
			Aliases: []string{"f", "-f"},
			Usage:   "run brainfuck code from file",
			Action:  actions.File,
		},
		{
			Name:    "interactive",
			Aliases: []string{"i", "-i"},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "simple",
					Usage:    "non fancy terminal",
					Required: false,
				},
			},
			Usage:       "run brainfuck code from stdin, just like terminal",
			Description: "exit by typing 'exit'. clear terminal by typing 'clear'",
			ArgsUsage:   "--simple for no fancy terminal",
			Action:      actions.Interactive,
		},
		{
			Name:    "read",
			Aliases: []string{"r", "-r"},
			Usage:   "run brainfuck code from args",
			Action:  actions.Arg,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
