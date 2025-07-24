package main

import (
	"fmt"
	"log"
	"os"
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
Start from {{ .from }}
{{ sh "hledger bal assets -b $from" }}
`

func loadTemplate() string {
	if len(os.Args) < 2 {
		log.Fatal("Please input a markdown template file")
	}
	filename := os.Args[1]
	dat, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to open the file", err)
	}
	return string(dat)
}

func main() {
	tmpl := loadTemplate()
	frontmatter, body, err := parseDocument(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	templateMeta, err := parseFrontmatter(frontmatter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(templateMeta)
	ctx, err := parseParams(templateMeta)
	if err != nil {
		log.Fatal(err)
	}
	data := render(body, ctx)
	output, err := templateMeta.OutputFile(ctx)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(output, data, 0644)
	if err != nil {
		panic(err)
	}
}
