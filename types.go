package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
