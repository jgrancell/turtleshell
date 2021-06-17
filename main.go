package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jgrancell/turtleshell/cli"
	"github.com/jgrancell/turtleshell/cmd"
	"github.com/jgrancell/turtleshell/configuration"
	"github.com/jgrancell/turtleshell/history"
	"github.com/jgrancell/turtleshell/variables"
)

var (
	version string = "1.0.0"
)

func main() {
	os.Exit(Terminal())
}

func Terminal() int {
	var err error
	var exitcode int

	shell := &cli.Cli{}
	shell.Version = version

	os.Setenv("SHELL", "/bin/turtleshell")

	shell.Configuration = configuration.Load()

	// Loading shell history
	shell.History, err = history.Load(shell.Configuration.HistoryFile)
	if err != nil {
		fmt.Println("Shell history initialization error:", err.Error())
		return 1
	}

	// Loading shell variables
	shell.Variables = variables.Load()

	// Catching all Interrupts and Kills and handling gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		for sig := range c {
			switch sig {
			case os.Interrupt:
				fmt.Println()
				fmt.Println(shell.Configuration.Prompt, " ")
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGINT:
				fmt.Println("Exiting turtleshell")
				os.Exit(1)
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		// Reloading on every command so that we get the most current configuration
		shell.Configuration.Reload()

		fmt.Print(shell.Configuration.Prompt, " ")

		var input string
		input, err = reader.ReadString('\n')

		if err != nil {
			fmt.Println("Command read error:", err.Error())
			exitcode = 127
			break
		} else {
			if input == "exit\n" {
				break
			} else if input != "\n" && input != "" {
				// Saving the command to our history
				err = shell.History.Append(strings.TrimSpace(input))
				if err != nil {
					fmt.Println("History append error:", err.Error())
					exitcode = 126
					break
				}
				exitcode, err = cmd.Parse(shell, input)
				if err != nil {
					fmt.Println(err.Error())
				}
				shell.LastStatus = exitcode
				continue
			}
		}
	}

	// Post run cleanup
	if err == nil {
		exitcode = 0
	}
	return exitcode
}
