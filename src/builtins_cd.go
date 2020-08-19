package main

import (
	"os"
	"strings"
)

func builtinsCd(args []string, conf Configuration) {
	target := strings.ReplaceAll(args[1], "~", conf.HomeDir)
	os.Chdir(target)
	setExitcode(0)
}
