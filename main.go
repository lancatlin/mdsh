package main

import (
	"errors"
	"fmt"
	"os"
)

var TMPL_EXAMPLE = `
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
	tmpl, err := loadTemplate()
	handleError(err)

	frontmatter, body, err := parseDocument(tmpl)
	handleError(err)

	templateMeta, err := parseFrontmatter(frontmatter)
	handleError(err)

	ctx, err := parseParams(templateMeta)
	handleError(err)

	data := render(body, ctx)
	output, err := templateMeta.OutputFile(ctx)
	handleError(err)

	err = os.WriteFile(output, data, 0644)
	handleError(err)

	fmt.Printf("Result saved to %s\n", output)
}

func loadTemplate() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Please input a markdown template file")
	}
	filename := os.Args[1]
	dat, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %w", err)
	}
	return string(dat), nil
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
