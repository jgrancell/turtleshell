package main

import (
	"fmt"
	"os"
	"strings"
)

func builtinsSet(args []string, conf Configuration) {
	if len(args) >= 3 {
		var value string
		if len(args) >= 4 {
			value = strings.Join(args[2:], " ")
			value = strings.Trim(value, `'"`)
		} else {
			value = args[2]
		}

		if args[1] == "exitcode" {
			setExitcode(98)
		} else {
			variable := strings.ToUpper(strings.Trim(args[1], "$"))
			os.Setenv(variable, value)
			setExitcode(0)
		}
	} else if len(args) < 3 {
		fmt.Println("Setting a variable must be done in the format of 'set $VARIABLE value'.")
		setExitcode(1)
	}
}
