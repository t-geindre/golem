package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed component.go.tmpl
var template string

const help = `Generates required code to retrieve the given component [type] on entities.
The generated code is stored in a "[type]_golem.go" file created in current working directory.

Usage:

	golem [package] [type]

Version:

	0.0.0`

func golem(stdout, stderr io.Writer, args []string) int {
	errLog := log.New(stderr, "", 0)

	// Get args
	if len(args) != 3 {
		_, _ = fmt.Fprintln(stdout, help)

		if len(args) == 2 && string(args[1]) == "help" {
			return 0
		}

		errLog.Println("Invalid arguments")
		return 1
	}

	pkg, tp := args[1], args[2]

	// Generate template
	r := strings.NewReplacer("{package}", pkg, "{type}", tp)
	code := r.Replace(template)

	// Write file
	e := os.WriteFile(fmt.Sprintf("%s_golem.go", strings.ToLower(tp)), []byte(code), 0644)
	if e != nil {
		errLog.Println(e)
		return 2
	}

	return 0
}

func main() {
	os.Exit(golem(os.Stdout, os.Stderr, os.Args))
}
