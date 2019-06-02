package main

import (
    "fmt"
    "os"
    "strings"
)

func builtins_get(args []string, conf Configuration) {
    if args[1] == "exitcode" {
        exitcode := os.Getenv("TURTLE_EXIT_CODE")
        fmt.Println(exitcode)
        os.Setenv("TURTLE_EXIT_CODE", "0")
    } else if args[1] == "all" {
        for _, e := range os.Environ() {
            if strings.HasPrefix(e, "TURTLE_") {
                fmt.Println(e)
            }
        }
    } else {
        variable := strings.ToUpper(strings.Trim(args[1], "$"))
        value := os.Getenv("TURTLE_" + variable)
        fmt.Println(value)
    }
}
