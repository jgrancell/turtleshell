package cmd

import (
	"fmt"

	"github.com/jgrancell/turtleshell/cli"
)

func Info(s *cli.Cli, args []string) (int, error) {
	fmt.Println("TurtleShell -- Version", s.Version)
	return 0, nil
}
