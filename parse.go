package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func parseDocument(doc string) (meta TemplateMeta, body string, err error) {
	frontmatter, body := separateDocument(doc)
	if frontmatter == "" {
		return TemplateMeta{}, body, nil
	}
	err = yaml.Unmarshal([]byte(frontmatter), &meta)
	if err != nil {
		err = fmt.Errorf("Failed to parse frontmatter: %w", err)
		return
	}
	return meta, body, nil
}

func separateDocument(doc string) (frontmatter string, body string) {
	doc = strings.TrimLeft(doc, " \n")
	sep := "---\n"
	if !strings.HasPrefix(doc, sep) {
		return "", doc
	}
	parts := strings.SplitN(doc, sep, 3)
	if len(parts) < 3 {
		return "", doc
	}
	frontmatter = parts[1]
	body = strings.TrimLeft(parts[2], "\n ")
	return
}
