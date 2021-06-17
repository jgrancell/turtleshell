package cmd

import (
	"fmt"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func Set(s *cli.Cli, components []string) (int, error) {
	var verbose bool

	for _, arg := range components {
		if arg == "--verbose" || arg == "-v" {
			verbose = true
		}
	}

	for _, arg := range components {
		if strings.Index(arg, "--") != 0 {
			a := strings.Split(arg, "=")
			if len(a) == 2 {
				if strings.Index(a[1], "$") == 0 {
					newvar := strings.TrimPrefix(a[1], "$")
					val, ok := s.Variables.Get(newvar)
					if ok {
						a[1] = val
					} else {
						a[1] = ""
					}
				}

				s.Variables.Set(a[0], a[1])
				if verbose {
					fmt.Println("Variable", a[0], "set to value", a[1])
				}
			}
		}
	}
	return 0, nil
}
