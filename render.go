package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func (ctx TemplateContext) run(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	for k, v := range ctx {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
	err := cmd.Run()
	if err != nil {
		return "", errors.New(stdout.String())
	}
	return stdout.String(), nil
}

func (ctx TemplateContext) raw(command string) string {
	out, err := ctx.run(command)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	return strings.TrimSpace(out)
}

func (ctx TemplateContext) codeBlock(command string) string {
	out, err := ctx.run(command)
	if err != nil {
		return fmt.Sprintf("> Failed to execute the command:\n> `%s`\n> %s\n", command, err.Error())
	}
	return fmt.Sprintf("\n```\n%s\n```\n", out)
}

func (ctx TemplateContext) inline(command string) string {
	out, err := ctx.run(command)
	if err != nil {
		return fmt.Sprintf("`Failed to execute the command: '%s'. %s`", command, err.Error())
	}
	return fmt.Sprintf("`%s`", strings.Trim(out, "\n "))
}

func (ctx TemplateContext) funcMap() template.FuncMap {
	return template.FuncMap{
		"sh":    ctx.inline,
		"shell": ctx.codeBlock,
		"raw":   ctx.raw,
	}
}

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
