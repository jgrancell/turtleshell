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
			os.Setenv("TURTLE_EXIT_CODE", args[1])
		} else {
			variable := strings.ToUpper(strings.Trim(args[1], "$"))
			os.Setenv("TURTLE_"+variable, value)
		}
	} else if len(args) < 3 {
		fmt.Println("Setting a variable must be done in the format of 'set $VARIABLE value'.")
	}
}
