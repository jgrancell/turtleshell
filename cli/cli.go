package cli

import (
	"fmt"
	"os"

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

func (c *Cli) Load(version string) error {
	var err error
	c.Version = version

	os.Setenv("SHELL", "/bin/turtleshell")

	c.Configuration = configuration.Load()

	// Loading shell history
	c.History, err = history.Load(c.Configuration.HistoryFile)
	if err != nil {
		return fmt.Errorf("shell history initialization error: %s", err.Error())
	}

	// Loading shell variables
	c.Variables = variables.Load()
	return nil
}
