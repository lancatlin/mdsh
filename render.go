package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"
)

func (ctx TemplateContext) render(tmpl string, output io.Writer) error {
	t, err := template.New("foo").Funcs(ctx.funcMap()).Parse(tmpl)
	if err != nil {
		return err
	}
	err = t.Execute(output, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ctx TemplateContext) funcMap() template.FuncMap {
	return template.FuncMap{
		"sh":    ctx.inline,
		"shell": ctx.codeBlock,
		"raw":   ctx.raw,
	}
}

func (ctx TemplateContext) OutputFile(meta TemplateMeta) (output io.WriteCloser, err error) {
	if meta.Output == "" {
		return os.Stdout, nil
	}
	tmpl, err := template.New("output").Funcs(ctx.funcMap()).Parse(meta.Output)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse output name: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute output name: %w", err)
	}

	filename := buf.String()

	fmt.Printf("Saving result to %s\n", filename)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to open the file: %w", err)
	}
	return file, nil
}
