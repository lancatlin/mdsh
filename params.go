package main

import (
	"flag"
	"fmt"
)

func parseParams(meta TemplateMeta) (ctx TemplateContext, err error) {
	result := make(map[string]*string)
	for key, spec := range meta.Params {
		result[key] = flag.String(key, spec.Default, key)
	}
	flag.Parse()
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
