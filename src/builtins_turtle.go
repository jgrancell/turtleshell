package main

import (
    "fmt"
    "os"
)

func builtins_turtle(args []string, conf Configuration) {
    fmt.Println("    _____    ____")
    fmt.Println("  /      \\  |  o | ")
    fmt.Println("  |        |/ __\\| ")
    fmt.Println("  |_________/     ")
    fmt.Println("  |_|_| |_|_|")
    fmt.Println()
    fmt.Println("TurtleShell -- Version", conf.Version)
    os.Setenv("TURTLE_EXIT_CODE", "0")
}
