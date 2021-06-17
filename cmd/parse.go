package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func Parse(s *cli.Cli, input string) (int, error) {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, " ")

	components := strings.Split(input, " ")
	if len(components) == 0 {
		return 0, nil
	}

	// Before we do any work we need to replace any variable components with true values
	for index, comp := range components {
		if strings.Index(comp, "$") == 0 {
			comp = strings.TrimPrefix(comp, "$")
			val, ok := s.Variables.Get(comp)
			if ok {
				components[index] = val
			} else {
				components[index] = ""
			}
		}

		if strings.Index(comp, "~") == 0 {
			components[index] = strings.Replace(comp, "~", s.Configuration.Home, 1)
		}
	}

	b := LoadBuiltins()
	builtin, err := b.Get(components[0])
	if err == nil {
		// The command is a shell builtin
		command := builtin.Command
		exitcode, err := command(s, components[1:])
		if err != nil {
			return exitcode, err
		}
		return exitcode, nil
	} else {
		// The command is not a builtin
		cmd := exec.Command(components[0], components[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			return 1, err
		}
	}

	return 0, nil
}
