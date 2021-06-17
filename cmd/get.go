package cmd

import (
	"fmt"

	"github.com/jgrancell/turtleshell/cli"
)

func Get(s *cli.Cli, components []string) (int, error) {
	if len(components) == 1 {
		val, ok := s.Variables.Get(components[0])
		if ok {
			fmt.Println(val)
			return 0, nil
		}
	}
	return 1, nil
}
