package cmd

import (
	"fmt"

	"github.com/jgrancell/turtleshell/cli"
)

func Status(s *cli.Cli, args []string) (int, error) {
	fmt.Println(s.LastStatus)
	return 0, nil
}
