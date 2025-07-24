package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var tmpl = `
---
output: "reports/${from}.md"
params:
  from:
    required: true
  file:
    default: "ledger.j"
---
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

func parseDocument(doc string) (frontmatter string, body string, err error) {
	doc = strings.TrimLeft(doc, " \n")
	sep := "---\n"
	if !strings.HasPrefix(doc, sep) {
		err = fmt.Errorf("Invalid document: frontmatter required.\n%s", doc)
		return
	}
	parts := strings.Split(doc, sep)
	if len(parts) < 3 {
		err = fmt.Errorf("Invalid document: frontmatter required.\n%s", doc)
		return
	}
	frontmatter = parts[1]
	body = parts[2]
	return
}

func main() {
	funcMap := template.FuncMap{
		"sh": sh,
	}
	frontmatter, body, err := parseDocument(tmpl)
	if err != nil {
		panic(err)
	}
	templateMeta, err := parseFrontmatter(frontmatter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("'%s'", templateMeta)
	t, err := template.New("foo").Funcs(funcMap).Parse(body)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "hledger is awesome")
}
