package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runnerCommand(args []string, conf Configuration) {

	biCheck := strings.Trim(args[0], `'`)

	if isBuiltin(biCheck) {
		builtins := NewBuiltins()
		command  := builtins.Commands[biCheck]
		command(args, conf)
	} else {
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stdin  = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
