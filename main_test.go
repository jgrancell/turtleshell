package main

import (
	"os"
	"testing"

	"github.com/jgrancell/turtleshell/cli"
	"github.com/jgrancell/turtleshell/configuration"
	"github.com/jgrancell/turtleshell/history"
	"github.com/jgrancell/turtleshell/variables"
)

func TestMain(t *testing.T) {

	// this should succeed
	exitcode := Terminal()
	if exitcode != 127 {
		t.Errorf("received terminal exitcode %d", exitcode)
	}

}

func TestParseInput(t *testing.T) {
	shell := &cli.Cli{}
	shell.Version = version

	os.Setenv("SHELL", "/bin/turtleshell")

	shell.Configuration = configuration.Load("")

	// Loading shell history
	var err error
	shell.History, err = history.Load("./testdata/history/history.txt")
	if err != nil {
		t.Errorf("Shell history initialization error: %s", err.Error())
	}

	// Loading shell variables
	shell.Variables = variables.Load()
	exitcode, ok := ParseInput(shell, "status")
	if !ok {
		t.Errorf("parse input for status failed, when it should have succeeded")
	}

	if exitcode != 0 {
		t.Errorf("exitcode for status should have been 0, was %d", exitcode)
	}
}
