package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"text/template"
)

func sh(ctx TemplateContext) func(string) string {
	return func(command string) string {
		cmd := exec.Command("sh", "-c", command)
		for k, v := range ctx {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
		out, err := cmd.Output()
		if err != nil {
			return fmt.Sprintf("> Failed to execute the command: %s\n> %s\n", cmd.Args, err.Error())
		}
		return fmt.Sprintf("\n```\n%s\n```\n", string(out))
	}
}

func render(tmpl string, ctx TemplateContext) []byte {
	funcMap := template.FuncMap{
		"sh": sh(ctx),
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
