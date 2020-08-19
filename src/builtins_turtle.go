package main

import (
	"fmt"
)

func builtinsTurtle(args []string, conf Configuration) {
	fmt.Println("    _____    ____")
	fmt.Println("  /      \\  |  o | ")
	fmt.Println("  |        |/ __\\| ")
	fmt.Println("  |_________/     ")
	fmt.Println("  |_|_| |_|_|")
	fmt.Println()
	fmt.Println("TurtleShell -- Version", conf.Version)
	setExitcode(0)
}
