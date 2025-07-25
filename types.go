package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

type TemplateMeta struct {
	Output string               `yaml:"output"`
	Params map[string]ParamSpec `yaml:"params"`
}

func (meta TemplateMeta) String() string {
	data, err := json.Marshal(meta)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (meta TemplateMeta) OutputFile(ctx TemplateContext) (output string, err error) {
	tmpl, err := template.New("output").Parse(meta.Output)
	if err != nil {
		return "", fmt.Errorf("Failed to parse output name: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to execute output name: %w", err)
	}
	return buf.String(), nil
}

type ParamSpec struct {
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
	Usage    string `yaml:"usage"`
}

type TemplateContext map[string]string

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
