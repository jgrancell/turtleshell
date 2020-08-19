package main

import (
	"fmt"
)

func builtinsExitcodes(args []string, conf Configuration) {
	fmt.Println("This is a list of all current exitcodes in TurtleShell.")
	fmt.Println("You can view the previous command's exitcode by running 'get exitcode'")
	fmt.Println("")
	fmt.Println("0    Successful.")
	fmt.Println("1    Error.")
	fmt.Println("98   You cannot manually set an exit code using the 'set' command.")
	fmt.Println("99   The specific command or options are not yet available.")
	fmt.Println("127  The target directory could not be found.")
	setExitcode(0)
}
