package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var tmpl = `
Hello this is the text:
{{ sh "hledger bal assets" }}
`

func sh(command string) string {
	args := strings.Split(command, " ")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return fmt.Sprintf("> Failed to execute the command:\n> %s\n", err.Error())
	}
	return fmt.Sprintf("\n```\n%s\n```\n", string(out))
}

func main() {
	funcMap := template.FuncMap{
		"sh": sh,
	}
	t, err := template.New("foo").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "hledger is awesome")
}
