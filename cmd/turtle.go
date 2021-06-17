package cmd

import (
	"fmt"

	"github.com/jgrancell/turtleshell/cli"
)

func Turtle(s *cli.Cli, args []string) (int, error) {
	fmt.Println("    _____    ____")
	fmt.Println("  /      \\  |  o | ")
	fmt.Println("  |        |/ __\\| ")
	fmt.Println("  |_________/     ")
	fmt.Println("  |_|_| |_|_|")
	fmt.Println()
	fmt.Println("TurtleShell -- Version", s.Version)
	return 0, nil
}
