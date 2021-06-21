package cmd

import (
	"testing"

	"github.com/jgrancell/turtleshell/cli"
	"github.com/jgrancell/turtleshell/variables"
)

func TestSetVariable(t *testing.T) {
	s := &cli.Cli{}
	s.Variables = variables.Load()

	if len(s.Variables.Variables) != 0 {
		t.Errorf("initial variable values should be 0, found %d", len(s.Variables.Variables))
	}

	// Testing a greenfield variable
	args := make([]string, 0)
	args = append(args, "foobar=baz")
	args = append(args, "--verbose")
	exitcode, err := Set(s, args)

	if exitcode != 0 || err != nil {
		t.Errorf("Set command should have returned 0 exitcode with no error, got %d %s", exitcode, err.Error())
	}

	if len(s.Variables.Variables) != 1 {
		t.Errorf("updated variable count should be 1, found %d", len(s.Variables.Variables))
	}

	entry := s.Variables.Variables[0]
	if entry.Name != "foobar" && entry.Value != "baz" {
		t.Errorf("variable entry name should be foobar and value should be baz, found %s %s", entry.Name, entry.Value)
	}

	// Testing replacing a variable value with a saved variable
	args = make([]string, 0)
	args = append(args, "fizzbuzz=$foobar")
	exitcode, err = Set(s, args)

	if exitcode != 0 || err != nil {
		t.Errorf("Set command should have returned 0 exitcode with no error, got %d %s", exitcode, err.Error())
	}

	if len(s.Variables.Variables) != 2 {
		t.Errorf("updated variable count should be 2, found %d", len(s.Variables.Variables))
	}

	entry = s.Variables.Variables[1]
	if entry.Name != "fizzbuzz" && entry.Value != "baz" {
		t.Errorf("variable entry name should be foobar and value should be baz, found %s %s", entry.Name, entry.Value)
	}

	// Testing replacing a variable value with a non-existant variable
	args = make([]string, 0)
	args = append(args, "fjords=$slartabartfast")
	exitcode, err = Set(s, args)

	if exitcode != 0 || err != nil {
		t.Errorf("Set command should have returned 0 exitcode with no error, got %d %s", exitcode, err.Error())
	}

	if len(s.Variables.Variables) != 3 {
		t.Errorf("updated variable count should be 3, found %d", len(s.Variables.Variables))
	}

	entry = s.Variables.Variables[2]
	if entry.Name != "fjords" && entry.Value != "" {
		t.Errorf("variable entry name should be fjors and value should be an empty string, found %s %s", entry.Name, entry.Value)
	}
}
