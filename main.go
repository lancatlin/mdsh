package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	doc, err := loadTemplate()
	handleError(err)

	templateMeta, body, err := parseDocument(doc)
	handleError(err)

	ctx, err := parseParams(templateMeta)
	handleError(err)

	output, err := ctx.OutputFile(templateMeta)
	handleError(err)
	defer output.Close()

	err = ctx.render(body, output)
	handleError(err)
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
