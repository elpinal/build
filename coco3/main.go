package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/elpinal/coco3-go-to-bed-by9/cli"
	"github.com/elpinal/coco3/config"
	"github.com/elpinal/color"
)

func ocamlVersion() string {
	s := os.Getenv("COCO3_OCAML_VERSION")
	if s == "" {
		s = "4.06.0"
	}
	return s
}

func main() {
	ocamlv := ocamlVersion()
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

				// OPAM for OCaml
				"OCAML_TOPLEVEL_PATH":  filepath.Join(home, ".opam/"+ocamlv+"/lib/toplevel"),
				"PERL5LIB":             filepath.Join(home, ".opam/"+ocamlv+"/lib/perl5"),
				"OPAMUTF8MSGS":         "1",
				"CAML_LD_LIBRARY_PATH": filepath.Join(home, ".opam/"+ocamlv+"/lib/stublibs"),

				// Rainy for Rain ML.
				"RAINY_RAIN_ML_URI": filepath.Join(home, "src/github.com/elpinal/rain-ml"),
				"RAINY_RAIN_VM_URI": filepath.Join(home, "src/github.com/elpinal/rain-vm"),
			},
			Paths: []string{
				filepath.Join(home, "bin"),
				filepath.Join(home, ".gvmn/go/current/bin"),
				filepath.Join(home, ".vvmn/vim/current/bin"),
				filepath.Join(home, ".aewo/bin"),
				filepath.Join(home, ".opam/"+ocamlv+"/bin"),
				filepath.Join(home, ".local/bin"),
				filepath.Join(home, ".psla/bin"),
				filepath.Join(home, ".cargo/bin"),
				filepath.Join(home, ".monumental-ruby/bin"),
				"/usr/local/share/dotnet",
				"/Library/TeX/texbin",
				"/usr/local/bin",
				"/usr/local/opt/coreutils/libexec/gnubin",
			},
			Extra: true,
		},
	}
	os.Exit(c.Run(os.Args[1:]))
}
