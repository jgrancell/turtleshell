package main

import (
	"fmt"
	"os"
	"os/user"
)

func builtinsGohome(args []string, conf Configuration) {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	if _, err := os.Stat(homeDir); err != nil {
		fmt.Println("The directory ", homeDir, "does not exist.")
		os.Setenv("TURTLE_EXIT_CODE", "127")
	} else {
		os.Chdir(homeDir)
		os.Setenv("TURTLE_EXIT_CODE", "0")
	}
}
