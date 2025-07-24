package main

import (
	"fmt"

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
