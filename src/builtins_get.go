package main

import (
    "fmt"
    "os"
)

func builtins_get(args []string, conf Configuration) {
    if args[1] == "exitcode" {
        exitcode := os.Getenv("TURTLE_EXIT_CODE")
        fmt.Println(exitcode)
        os.Setenv("TURTLE_EXIT_CODE", "0")
    }
}
