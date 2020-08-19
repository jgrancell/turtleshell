package main

import (
	"fmt"
	"os"
	"strings"
)

func builtinsEcho(args []string, conf Configuration) {
	// For the echo builtin, there's potentially up to 4 arguments
	// args[0] == the echo command itself
	// args[1] == the variable of string we're echoing
	// args[2] == the append or overwrite symbol
	// args[3] == the append or overwrite target
	if len(args) > 1 {
		string := args[1]
		// Determining if we're echoing into a file
		if len(args) == 4 {
			//targetFile := args[3]
			// Determining if we're appending or overwriting
			if args[2] == ">" {
				// Overwriting our target file
				//overwrite := true
				setExitcode(99)
			} else {
				// Appending to our target file
				//overwrite := false
				setExitcode(99)
			}
		} else {
			// We're either printing just a string, or a variable
			if strings.HasPrefix(string, "$") {
				// We're looking for the value of a variable
				if _, ok := variables[string]; ok {
					fmt.Println(variables[string])
					setExitcode(0)
				} else {
					variable := strings.ToUpper(strings.Trim(string, "$"))
					if value, ok := os.LookupEnv(variable); ok {
						fmt.Println(value)
						setExitcode(0)
					} else {
						setExitcode(1)
					}
				}
			} else {
				fmt.Println(strings.Trim("\"", string))
				setExitcode(0)
			}
		}
	} else {
		setExitcode(1)
	}
}
