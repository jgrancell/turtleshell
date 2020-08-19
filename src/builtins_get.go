package main

import (
	"fmt"
	"os"
	"strings"
)

func builtinsGet(args []string, conf Configuration) {
	if args[1] == "exitcode" {
		fmt.Println(exitCode)
		setExitcode(0)
	} else if args[1] == "all" {
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
		setExitcode(0)
	} else {
		variable := strings.ToUpper(strings.Trim(args[1], "$"))
		value := os.Getenv(variable)
		fmt.Println(value)
		setExitcode(0)
	}
}
