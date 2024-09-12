package golemgen

import (
	_ "embed"
	"strings"
)

//go:embed component.go.tmpl
var cmpTemplate string

type Templater struct {
}

func NewTemplater() *Templater {
	return &Templater{}
}

func (t *Templater) Apply(pkg, tp string) string {
	r := strings.NewReplacer("{package}", pkg, "{type}", tp)
	return r.Replace(cmpTemplate)
}
