package cmd

import (
	"fmt"
	"os"

	"github.com/jgrancell/turtleshell/cli"
)

func Run(s *cli.Cli, components []string) (int, error) {
	if len(components) == 1 {
		target := components[0]

		if info, err := os.Stat(target); err == nil {
			mode := info.Mode()
			if mode&0111 != 0 {
				// TODO: add script runner
				return 0, fmt.Errorf("todo feature not yet implemented to run target script file %s", target)
			} else {
				return 1, fmt.Errorf("the target script file %s is not executable", target)
			}
		} else {
			return 1, err
		}
	} else {
		return 1, fmt.Errorf("the run command only takes a single argument, the file to be run")
	}
}
