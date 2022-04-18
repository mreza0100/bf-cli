package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mreza0100/brainfuck"
	"github.com/urfave/cli"
)

type Actions struct{}

func (_ *Actions) exec(reader io.Reader) {
	bf := brainfuck.New()
	bf.Entry(reader)
}

func (a *Actions) File(c *cli.Context) error {
	_, args := isInArgs(c.Args(), append(c.Command.Aliases, c.Command.Name))
	if len(args) == 0 {
		return cli.NewExitError("missing file name/path", 1)
	}

	reader, err := os.Open(args[0])
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	a.exec(reader)

	return nil
}

func (a *Actions) Arg(c *cli.Context) error {
	_, args := isInArgs(c.Args(), append(c.Command.Aliases, c.Command.Name))

	allInstructions := strings.Join(args, "")
	reader := strings.NewReader(allInstructions)

	a.exec(reader)

	return nil
}

func (a *Actions) Interactive(c *cli.Context) error {
	clearTerminal()
	isSimpleMode := c.Bool("simple")

	if isSimpleMode {
		a.exec(os.Stdin)
		return nil
	}

	var input string
Runner:
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&input); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		switch input {
		case "exit":
			break Runner
		case "clear":
			clearTerminal()
			continue Runner
		default:
			a.exec(strings.NewReader(input))
		}
	}

	return nil
}
