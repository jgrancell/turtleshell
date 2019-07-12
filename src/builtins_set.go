package main

import (
	"fmt"
	"os"
	"strings"
)

func builtinsSet(args []string, conf Configuration) {
	if len(args) == 3 {
		if args[1] == "exitcode" {
			os.Setenv("TURTLE_EXIT_CODE", args[1])
		} else {
			variable := strings.ToUpper(strings.Trim(args[1], "$"))
			os.Setenv("TURTLE_"+variable, args[2])
		}
	} else if len(args) < 3 {
		fmt.Println("Setting a variable must be done in the format of 'set $VARIABLE value'.")
	}
}
