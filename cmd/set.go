package cmd

import (
	"fmt"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
	"github.com/jgrancell/turtleshell/utils"
)

func Set(s *cli.Cli, components []string) (int, error) {
	// TODO: Add support for --export to environment
	verbose := utils.HasVerbose(components)

	args := utils.UnsetFlags(components)

	for _, arg := range args {
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
	return 0, nil
}
