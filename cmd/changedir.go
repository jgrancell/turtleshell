package cmd

import (
	"os"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func Changedir(s *cli.Cli, components []string) (int, error) {
	target := strings.ReplaceAll(components[0], "~", s.Configuration.Home)
	if err := os.Chdir(target); err != nil {
		return 1, err
	}
	return 0, nil
}
