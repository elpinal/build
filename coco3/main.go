package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/elpinal/coco3/cli"
	"github.com/elpinal/coco3/config"
	"github.com/elpinal/color"
)

func main() {
	home := os.Getenv("HOME")
	c := cli.CLI{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,

		Config: config.Config{
			PromptTmpl: template.Must(template.New("prompt").Parse(fmt.Sprintf("\n%s\n∆ ", color.Wrap("{{.WD}}", color.Yellow)))),
			Env: map[string]string{
				"EDITOR":   "vim",
				"PAGER":    "less",
				"GOPATH":   home,
				"GHQ_ROOT": filepath.Join(home, "src"),
			},
			StartUpCommand: []byte(`
			setpath (
				~/bin
				~/.gvmn/go/current/bin
				~/.vvmn/vim/current/bin
				~/.aewo/bin
				~/.opam/system/bin
				~/.local/bin
                                ~/.psla/bin
				~/.cargo/bin
				~/.monumental-ruby/bin
				/Library/TeX/texbin
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
				{"screen", "let LANG UTF-8 in screen"},
			},
			Extra: true,
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
