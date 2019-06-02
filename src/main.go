package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "strings"
    "strconv"
)

func main() {

    conf := loadConfiguration()

    if args := os.Args[1:]; len(args) > 0 {
        execInput(strings.Join(args, " "), conf)
        os.Exit(0)
    }

    sigCatch(conf)

    reader := bufio.NewReader(os.Stdin)
    for {
        conf := loadConfiguration()

        // Provide the shell prompt
        fmt.Print(conf.Prompt, " ")

        // Read the keyboard input
        input, err := reader.ReadString('\n')

        if err != nil {
            fmt.Println()
        } else {
            if input != "\n" && input != ""  {
                // Execute the user input
                execInput(input, conf)
            }    
        }
    }
}

func sigCatch(conf Configuration) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for sig := range c {
            if sig == os.Interrupt {
                fmt.Println()
                fmt.Print(conf.Prompt, " ")
            }
        }
    }()
}

func execInput(input string, conf Configuration) {
    // Remove the newline character
    input = strings.TrimSuffix(input, "\n")

    // Separate the command from the arguments
    args := strings.Split(input, " ")

    // Determining if we're receiving a discrete exit command
    if args[0] == "exit" {
        code := 0
        if len(args) == 2 {
            if parsedCode, err := strconv.Atoi(args[1]); err == nil {
                code = parsedCode
            } else {
                code = 155
            }
        }
        os.Exit(code)
    } else {

        // Determining if we're running a builtin command
        if isBuiltin(args[0]) {
            builtins := NewBuiltins()
            command := builtins.Commands[args[0]]
            command(args, conf)
        } else {
            // Determining what sort of command we're running
            info := runner_identify(args)

            // Running a script
            if info["runner"] == "file" {
                runner_scripts(args)
            } else {
                // Running an external command
                // Prepare the command to execute
                cmd := exec.Command(args[0], args[1:]...)

                // Attach to the output devices
                cmd.Stderr = os.Stderr
                cmd.Stdout = os.Stdout

                // Run the command and return the output
                if err := cmd.Run(); err != nil {
                    fmt.Fprintln(os.Stderr, err)
                }
            }
        }
    }
}
