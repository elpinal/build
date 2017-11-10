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

		Config: &config.Config{
			PromptTmpl: template.Must(template.New("prompt").Parse(fmt.Sprintf("\n%s\n∆ ", color.Wrap("{{.WD}}", color.Yellow)))),
			Env: map[string]string{
				"EDITOR":   "vim",
				"PAGER":    "less",
				"GOPATH":   home,
				"GHQ_ROOT": filepath.Join(home, "src"),
			},
			Paths: []string{
				filepath.Join(home, "bin"),
				filepath.Join(home, ".gvmn/go/current/bin"),
				filepath.Join(home, ".vvmn/vim/current/bin"),
				filepath.Join(home, ".aewo/bin"),
				filepath.Join(home, ".opam/system/bin"),
				filepath.Join(home, ".local/bin"),
				filepath.Join(home, ".psla/bin"),
				filepath.Join(home, ".cargo/bin"),
				filepath.Join(home, ".monumental-ruby/bin"),
				"/Library/TeX/texbin",
				"/usr/local/bin",
				"/usr/local/opt/coreutils/libexec/gnubin",
			},
			Extra: true,
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
