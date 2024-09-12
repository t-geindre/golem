package main

import (
	"github.com/t-geindre/golem/internal/golemgen"
	"log"
	"os"
	"strings"
)

func main() {
	errLog := log.New(os.Stderr, "", 0)

	if len(os.Args) < 2 {
		errLog.Println("file to parse missing")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		errLog.Println("too many arguments")
		os.Exit(1)
	}

	file := string(os.Args[1])

	parser := golemgen.NewParser()
	e := parser.Parse(file)

	if e != nil {
		errLog.Println(e)
		os.Exit(2)
	}

	templater := golemgen.NewTemplater()
	code := templater.Apply(parser.GetPackage(), parser.GetType())

	i := strings.LastIndex(file, ".go")
	out := file[:i] + strings.Replace(file[i:], ".go", "_golem.go", 1)

	e = os.WriteFile(out, []byte(code), 0644)

	if e != nil {
		errLog.Println(e)
		os.Exit(3)
	}
}
