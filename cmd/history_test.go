package cmd

import (
	"testing"

	"github.com/jgrancell/turtleshell/cli"
)

func TestHistory(t *testing.T) {
	s := &cli.Cli{}
	s.Load("0.0.1-test")

	components := make([]string, 0)
	exitcode, _ := History(s, components)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}

	components = make([]string, 0)
	components = append(components, "foobar")
	exitcode, _ = History(s, components)
	if exitcode != 0 {
		t.Errorf("expected a success exitcode 0, got %d", exitcode)
	}
}
