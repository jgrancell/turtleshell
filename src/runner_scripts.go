package main

import (
	"fmt"
	"os"
)

func runnerScripts(args []string) {
	// Determining if a file target was passed
	if len(args) > 1 {
		// Determining if the file target exists
		target := args[1]

		if info, err := os.Stat(target); err == nil {
			// Determining if we have permissions to execute the file
			mode := info.Mode()

			if mode&0111 != 0 {
				fmt.Println("Executing script", target)
				setExitcode(0)
			} else {
				fmt.Println("\033[31m"+target, "is not executable.\033[0m")
				setExitcode(1)
			}
		} else {
			fmt.Println("\033[31mThe file", target, "does not exist.\033[0m")
			setExitcode(1)
		}
	} else {
		fmt.Println("\033[31mNo script was specified to run.\033[0m")
		setExitcode(1)
	}
}
