package main

import (
	"flag"
	"fmt"
	"os"
)

func parseParams(meta TemplateMeta) (ctx TemplateContext, err error) {
	result := make(map[string]*string)
	fs := flag.NewFlagSet("params", flag.ExitOnError)
	for key, spec := range meta.Params {
		result[key] = fs.String(key, spec.Default, key)
	}
	fs.Parse(os.Args[2:])
	ctx = make(TemplateContext)
	for key, spec := range meta.Params {
		if spec.Required && *result[key] == "" {
			err = fmt.Errorf("parameter '%s' required but not provided", key)
			return
		}
		ctx[key] = *result[key]
	}
	return
}
