package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func Parse(s *cli.Cli, input string) (int, error) {
	var exitcode int
	var err error

	components := ParseInput(s, input)

	if len(components) == 0 {
		return 0, nil
	}

	b := LoadBuiltins()
	builtin, err := b.Get(components[0])
	if err == nil {
		// The command is a shell builtin
		command := builtin.Command
		exitcode, err = command(s, components[1:])
	} else {
		exitcode, err = RunBinary(components)
	}

	if exitcode != 0 {
		return exitcode, err
	}
	return exitcode, nil
}

func ParseInput(s *cli.Cli, input string) []string {
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")

	components := strings.Split(input, " ")
	components = ReplaceVariablesWithValues(s, components)
	return components
}

func ReplaceVariablesWithValues(s *cli.Cli, components []string) []string {
	for index, comp := range components {
		if strings.Index(comp, "$") == 0 {
			comp = strings.TrimPrefix(comp, "$")
			components[index], _ = s.Variables.Get(comp)
		}

		if strings.Index(comp, "~") == 0 {
			components[index] = strings.Replace(comp, "~", s.Configuration.Home, 1)
		}
	}
	return components
}

func RunBinary(components []string) (int, error) {
	// The command is not a builtin
	cmd := exec.Command(components[0], components[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return 1, err
	}
	return 0, nil
}
