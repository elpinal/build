package main

import (
	"os"

	"github.com/elpinal/coco3/cli"
	"github.com/elpinal/coco3/config"
)

func main() {
	c := cli.CLI{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,

		Config: config.Config{
			Prompt: "∆ ",
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
