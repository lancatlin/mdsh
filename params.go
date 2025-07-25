package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func parseParams(meta TemplateMeta) (ctx TemplateContext, err error) {
	result := make(map[string]*string)
	fs := flag.NewFlagSet("params", flag.ExitOnError)
	for key, spec := range meta.Params {
		result[key] = fs.String(key, spec.Default, spec.Usage)
	}
	fs.Parse(os.Args[2:])
	ctx = make(TemplateContext)
	message := ""
	for key, spec := range meta.Params {
		if spec.Required && *result[key] == "" {
			message += fmt.Sprintf("parameter '%s' required but not provided\n", key)
			continue
		}
		ctx[key] = *result[key]
	}
	if message != "" {
		message += fmt.Sprintf("Use 'mdsh %s --help' to see usage", os.Args[1])
		return nil, errors.New(message)
	}
	return
}
