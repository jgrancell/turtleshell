package cmd

import (
	"testing"

	"github.com/jgrancell/turtleshell/cli"
)

func TestRun(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	components := make([]string, 0)
	components = append(components, "../testdata/cmd/run/script.tsh")
	exitcode, _ := Run(s, components)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, "../testdata/cmd/run/script-nonexecute.tsh")
	exitcode, _ = Run(s, components)
	if exitcode != 1 {
		t.Errorf("expected error 1 from nonexecutable script file, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, "../testdata/cmd/run/nonexistant.tsh")
	exitcode, _ = Run(s, components)
	if exitcode != 1 {
		t.Errorf("expected missing script file error 1, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, "/target/one")
	components = append(components, "/target/two")
	exitcode, _ = Run(s, components)
	if exitcode != 1 {
		t.Errorf("expected failure code 1 due to too many args, got %d", exitcode)
	}
}
