package main

import (
	"fmt"
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
	fmt.Printf("'%s'", templateMeta)
	render(body, map[string]string{
		"from": "2025-07-15",
	})
}
