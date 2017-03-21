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
			StartUpCommand: []byte(`
			setenv EDITOR vim PAGER less
			setpath /usr/local/bin /usr/local/opt/coreutils/libexec/gnubin
			`),
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
