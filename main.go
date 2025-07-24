package main

import (
	"fmt"
	"log"
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

func main() {
	frontmatter, body, err := parseDocument(tmpl)
	if err != nil {
		panic(err)
	}
	templateMeta, err := parseFrontmatter(frontmatter)
	if err != nil {
		panic(err)
	}
	fmt.Println(templateMeta)
	ctx, err := parseParams(templateMeta)
	if err != nil {
		log.Fatal(err)
	}
	render(body, ctx)
}
