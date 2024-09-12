package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGolemArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "no args", args: []string{"cmd"}, want: 1},
		{name: "not enough", args: []string{"cmd", "foo"}, want: 1},
		{name: "too many", args: []string{"cmd", "foo", "bar", "foo"}, want: 1},
		{name: "help", args: []string{"cmd", "help"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bOut, bErr bytes.Buffer
			stdout, stderr := bufio.NewWriter(&bOut), bufio.NewWriter(&bErr)
			if got := golem(stdout, stderr, tt.args); got != tt.want {
				t.Errorf("golem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGolemTemplating(t *testing.T) {
	pkg := "mypackage"
	tp := "Mytype"
	file := fmt.Sprintf("%s_golem.go", strings.ToLower(tp))

	var bOut, bErr bytes.Buffer
	stdout, stderr := bufio.NewWriter(&bOut), bufio.NewWriter(&bErr)

	if got := golem(stdout, stderr, []string{"cmd", pkg, tp}); got != 0 {
		t.Errorf("golem() = %v, want %v", got, 0)
	}

	if _, err := os.Stat(file); err != nil {
		t.Fatalf("golem() file %s not created", file)
	}

	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("golem() file %s not readable", file)
	}

	expected, _ := os.ReadFile("testdata/expected.do.tmpl")

	if string(content) != string(expected) {
		t.Fatalf("golem() invalid content generated")
	}

	os.Remove(file)
}
