package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func parseFrontmatter(input string) (meta TemplateMeta, err error) {
	err = yaml.Unmarshal([]byte(input), &meta)
	if err != nil {
		err = fmt.Errorf("Failed to parse frontmatter: %w", err)
		return
	}
	return
}

func parseDocument(doc string) (frontmatter string, body string, err error) {
	doc = strings.TrimLeft(doc, " \n")
	sep := "---\n"
	if !strings.HasPrefix(doc, sep) {
		err = fmt.Errorf("Invalid document: frontmatter required.\n%s", doc)
		return
	}
	parts := strings.SplitN(doc, sep, 3)
	if len(parts) < 3 {
		err = fmt.Errorf("Invalid document: frontmatter required.\n%s", doc)
		return
	}
	frontmatter = parts[1]
	body = strings.TrimLeft(parts[2], "\n ")
	return
}
