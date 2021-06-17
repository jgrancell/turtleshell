package cmd

import (
	"os"

	"github.com/jgrancell/turtleshell/cli"
)

func Home(s *cli.Cli, args []string) (int, error) {
	if err := os.Chdir(s.Configuration.Home); err != nil {
		return 1, err
	}
	return 0, nil
}
