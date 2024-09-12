package golemgen

import (
	"fmt"
	"os"
	"regexp"
)

type Parser struct {
	pkg string // Package name
	tp  string // Component type
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(file string) error {
	src, e := os.ReadFile(file)
	if e != nil {
		return e
	}

	pkgReg := regexp.MustCompile(`package\s+(\w+)`)
	pkg := pkgReg.FindStringSubmatch(string(src))

	if len(pkg) == 0 {
		return fmt.Errorf("package not found")
	}

	if len(pkg) > 2 {
		return fmt.Errorf("multiple package declarations")
	}

	p.pkg = pkg[1]

	tpReg := regexp.MustCompile(`type\s+(\w+)\s+struct`)
	tp := tpReg.FindStringSubmatch(string(src))

	if len(tp) == 0 {
		return fmt.Errorf("component type not found")
	}

	if len(tp) > 2 {
		return fmt.Errorf("multiple component type declarations")
	}

	p.tp = tp[1]

	return nil
}

func (p *Parser) GetPackage() string {
	return p.pkg
}

func (p *Parser) GetType() string {
	return p.tp
}
