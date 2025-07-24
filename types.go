package main

import "encoding/json"

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

type ParamSpec struct {
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
}

type TemplateContext map[string]string
