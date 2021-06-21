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
)

var (
	version string = "1.0.0"
)

func main() {
	os.Exit(Terminal())
}

func Terminal() int {
	shell := &cli.Cli{}
	err := shell.Load(version)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

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
			default:
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

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Exiting turtleshell:", err.Error())
			return 127

		}

		exitcode, ok := ParseInput(shell, input)
		if !ok {
			return exitcode
		}
	}
}

func ParseInput(shell *cli.Cli, input string) (int, bool) {

	if input == "exit\n" {
		return 0, false
	} else if input != "\n" && input != "" {
		// Saving the command to our history
		err := shell.History.Append(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("History append error:", err.Error())
			return 126, false
		}
		exitcode, err := cmd.Parse(shell, input)
		if err != nil {
			fmt.Println(err.Error())
		}
		shell.LastStatus = exitcode

		return 0, true
	}
	shell.LastStatus = 1
	return 1, true
}
