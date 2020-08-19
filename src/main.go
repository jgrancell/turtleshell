package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var exitCode int
var history []string
var variables map[string]string

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
			if input != "\n" && input != "" {
				// Execute the user input
				historySave(input, conf)
				execInput(input, conf)

				// Save it to active shell history
				history = append(history, strings.TrimSpace(input))
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
		// Determining what sort of command we're running
		info := runnerIdentify(args)

		// Running a script
		if info["runner"] == "file" {
			runnerScripts(args)
		} else if info["runner"] == "command" {
			runnerCommand(args[1:], conf)
		} else {
			runnerCommand(args, conf)
		}
	}
}
