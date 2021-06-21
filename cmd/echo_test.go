package cmd

import (
	"testing"

	"github.com/jgrancell/turtleshell/cli"
)

func TestEcho(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	components := make([]string, 0)
	components = append(components, "Hello")
	components = append(components, "World")
	exitcode, _ := Echo(s, components)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, "Hello")
	components = append(components, ">>")
	components = append(components, "World")
	exitcode, _ = Echo(s, components)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, ">")
	components = append(components, "Hello")
	components = append(components, "World")
	exitcode, _ = Echo(s, components)
	if exitcode != 1 {
		t.Errorf("expected a vailure exitcode 1, got %d", exitcode)
	}
}
