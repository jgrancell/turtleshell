package main

import (
	"fmt"
	"strconv"
	"strings"
)

func builtinsHistory(args []string, conf Configuration) {
	if len(args) > 1 {
		lookupValue := historyParseValue(args[1:])
		fmt.Println("Showing all history for the active shell that matches", lookupValue)
		matches := historyGetMatches(lookupValue, history)
		for _, match := range matches {
			fmt.Println("  ", match)
		}
		setExitcode(0)
	} else {
		fmt.Println("For complete shell history see $HOME/.turtle_history.")
		for _, hist := range history {
			fmt.Println("  ", hist)
		}
		setExitcode(0)
	}
}

func builtinsRedo(args []string, conf Configuration) {
	var matches []string
	if len(args) > 1 {
		lookupValue := historyParseValue(args[1:])
		matches = historyGetMatches(lookupValue, history)
	} else {
		matches = history
	}

	fmt.Println("Type the number associated with the history item you wish to rerun:")

	for int, match := range matches {
		fmt.Println ("  "+strconv.Itoa(int)+": '"+match+"'")
	}
	var choice int
	fmt.Scanln(&choice)
	input := matches[choice]
	historySave(input, conf)
	execInput(input, conf)
	history = append(history, strings.TrimSpace(input))
}

func historyParseValue(args []string) string {
	var value string
	// cleaning our arguments and turning them into one value
	if len(args) >= 2 {
		value = strings.Join(args, " ")
		value = strings.Trim(value, `'"`)
	} else {
		value = args[0]
	}
	return value
}

func historyGetMatches(value string, history []string) []string {
	var matches []string
	for _, hist := range history {
		if strings.Contains(hist, value) {
			matches = append(matches, hist)
		}
	}

	return matches
}
