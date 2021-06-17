package cli

import (
	"github.com/jgrancell/turtleshell/configuration"
	"github.com/jgrancell/turtleshell/history"
	"github.com/jgrancell/turtleshell/variables"
)

type Cli struct {
	Configuration *configuration.Configuration
	LastStatus    int
	History       *history.History
	Variables     *variables.Variables
	Version       string
}
