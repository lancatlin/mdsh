package main

import (
	"bytes"
	"text/template"
)

func render(tmpl string, ctx TemplateContext) []byte {
	funcMap := template.FuncMap{
		"sh":    ctx.inline,
		"shell": ctx.codeBlock,
	}
	t, err := template.New("foo").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, ctx)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
