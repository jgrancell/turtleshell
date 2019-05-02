package main

import (
    "fmt"
    "os"
)

func builtins_exitcodes(args []string, conf Configuration) {
    fmt.Println("This is a list of all current exitcodes in TurtleShell.")
    fmt.Println("You can view the previous command's exitcode by running 'get exitcode'")
    fmt.Println("")
    fmt.Println("0    Successful.")
    fmt.Println("1    Error.")
    fmt.Println("127  The target directory could not be found.")
    os.Setenv("TURTLE_EXIT_CODE", "0")
}
