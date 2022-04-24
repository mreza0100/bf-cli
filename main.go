package main

import (
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	app.Name = "brainfuck-cli"
	app.Usage = "brainfuck-cli"
	app.Version = "0.1.8"
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
			Name:    "interactive - REPL",
			Aliases: []string{"i", "-i"},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "simple",
					Usage:    "non fancy terminal",
					Required: false,
				},
			},
			Usage:       "run brainfuck code in REPL mode",
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
