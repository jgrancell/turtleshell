package main

import (
	"os"
	"strings"
)

func builtins_cd(args []string, conf Configuration) {
	target := strings.ReplaceAll(args[1], "~", conf.HomeDir)
	os.Chdir(target)
	os.Setenv("TURTLE_EXIT_CODE", "0")
}
