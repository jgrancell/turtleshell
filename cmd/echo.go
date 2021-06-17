package cmd

import (
	"fmt"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func Echo(s *cli.Cli, components []string) (int, error) {
	// TODO: add the ability to echo output into files
	var redirect bool
	var content []string
	var targets []string
	for index, comp := range components {
		if comp == ">" || comp == ">>" {
			if index == 0 {
				return 1, fmt.Errorf("invalid echo command: no content provided for append or overwrite")
			}

			if index >= 1 && index < (len(components)-1) {
				redirect = true
				content = components[:index]
				targets = components[index+1:]
				break
			}
		}
	}

	if redirect {
		// TODO: Do the actuall file creation/appending/overwriting at some point
		fmt.Println("This feature is not yet available, but in the future this would send the content provided to your chosen targets.")
		fmt.Println("Content:", content)
		fmt.Println("Targets:", targets)
	} else {
		fmt.Println(strings.Trim(strings.Join(components[:], " "), " "))
	}
	return 0, nil
}
