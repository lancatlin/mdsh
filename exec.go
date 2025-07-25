package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
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
