package cmd

import (
	"testing"

	"github.com/jgrancell/turtleshell/cli"
)

func TestParseInput(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	input := "   foo bar-baz "
	out := ParseInputs(s, input)

	if len(out) != 2 {
		t.Errorf("expected 2 processed components, got %d", len(out))
	}

	if out[0] != "foo" && out[1] != "bar-baz" {
		t.Errorf("expected arguments to be foo, bar-baz. got %s, %s", out[0], out[1])
	}

	s.Variables.Set("FIZZBUZZ", "fjords")
	input = "    slartibartfast created the $FIZZBUZZ"
	out = ParseInputs(s, input)
	if len(out) != 4 {
		t.Errorf("expected 4 processed components, got %d", len(out))
	}
	if out[3] != "fjords" {
		t.Errorf("expected component 4 to be 'fjords', got %s", out[3])
	}
}

func TestBinaryCommand(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	args := make([]string, 0)
	args = append(args, "test")
	args = append(args, "-f")
	args = append(args, "/tmp/foobar.baz")
	exitcode, _ := RunBinary(args)

	if exitcode != 1 {
		t.Errorf("expected an error exitcode 1, got %d", exitcode)
	}

	args[1] = "-d"
	args[2] = "/tmp"
	exitcode, _ = RunBinary(args)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}
}

func TestParse(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	args := "test -d /tmp"
	exitcode, _ := Parse(s, args)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}
}
