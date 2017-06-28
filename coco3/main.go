package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/elpinal/coco3/cli"
	"github.com/elpinal/coco3/config"
	"github.com/elpinal/color"
)

func main() {
	c := cli.CLI{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,

		Config: config.Config{
			PromptTmpl: template.Must(template.New("prompt").Parse(fmt.Sprintf("\n%s\n∆ ", color.Wrap("{{.WD}}", color.Yellow)))),
			StartUpCommand: []byte(`
			setenv (
				EDITOR   vim
				PAGER    less
				GOPATH   ~
				GHQ_ROOT ~/src
			)
			setpath (
				~/bin
				~/.gvmn/go/current/bin
				~/.vvmn/vim/current/bin
				~/.aewo/bin
				~/.opam/system/bin
				/usr/local/bin
				/usr/local/opt/coreutils/libexec/gnubin
			)
			`),
			Alias: [][2]string{
				{"..", "cd .."},
				{"ls", "ls --show-control-chars --color=auto"},
				{"la", "ls -a"},
				{"ll", "ls -l"},
				{"lla", "ls -la"},
				{"v", "vim"},
				{"g", "git"},
				{"screen", "let LANG '' in screen"},
			},
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
