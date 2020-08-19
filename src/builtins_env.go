package main

import (
	"fmt"
	"os"
)

func builtinsEnv(args []string, conf Configuration) {
	value := os.Environ()
	fmt.Println(value)
	os.Setenv("TURTLE_EXIT_CODE", "0")
}
