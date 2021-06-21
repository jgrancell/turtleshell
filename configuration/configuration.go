package configuration

import (
	"encoding/json"
	"os"
	"os/user"
	"strings"

	"github.com/jgrancell/turtleshell/utils"
)

// Configuration information for the application
type Configuration struct {
	PS1          string `json:"ps1"`
	Prompt       string
	User         string
	HistoryFile  string
	HistoryLimit int `json:"history_limit"`
	Home         string
	RcFile       string
	Version      string
}

func Load(rcfile string) *Configuration {
	c := &Configuration{}
	c.user()

	if rcfile == "" {
		c.RcFile = c.Home + "/.turtlerc"
	} else {
		c.RcFile = rcfile
	}

	file, err := utils.OpenFile(c.RcFile)
	if err == nil {
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(c)
	} else {
		_ = utils.ValidateOrCreateFile(c.RcFile)
	}
	_ = file.Close()

	if c.PS1 == "" {
		c.PS1 = "[{user}@{hostname}]: >"
	}

	if c.HistoryLimit < 1 {
		c.HistoryLimit = 20
	}
	c.generatePrompt()
	c.HistoryFile = c.Home + "/.turtle_history"

	return c
}

func (c *Configuration) Reload() {
	file, err := utils.OpenFile(c.RcFile)
	if err == nil {
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(c)
	}
	_ = file.Close()
	c.generatePrompt()
}

func (c *Configuration) generatePrompt() {
	c.Prompt = c.PS1

	// Populateable shortcodes
	cwd, _ := os.Getwd()
	cwdTilded := strings.ReplaceAll(cwd, c.Home, "~")
	hostname, _ := os.Hostname()

	replacements := map[string]string{
		"{cwd}":       cwd,
		"{cwd:tilde}": cwdTilded,
		"{user}":      c.User,
		"{hostname}":  hostname,

		"{color:red}":          "\033[31m",
		"{color:red:bold}":     "\033[31;1m",
		"{color:green}":        "\033[32m",
		"{color:green:bold}":   "\033[32;1m",
		"{color:yellow}":       "\033[33m",
		"{color:yellow:bold}":  "\033[33;1m",
		"{color:blue}":         "\033[34m",
		"{color:blue:bold}":    "\033[34;1m",
		"{color:magenta}":      "\033[35m",
		"{color:magenta:bold}": "\033[35;1m",
		"{color:cyan}":         "\033[36m",
		"{color:cyan:bold}":    "\033[36;1m",
		"{color:white}":        "\033[37m",
		"{color:white:bold}":   "\033[37;1m",
		"{endcolor}":           "\033[0m",

		c.Home: "~",
	}

	for k, v := range replacements {
		if strings.Contains(c.Prompt, k) {
			c.Prompt = strings.ReplaceAll(c.Prompt, k, v)
		}
	}
}

func (c *Configuration) user() {
	usr, _ := user.Current()

	c.User = usr.Username
	_ = os.Setenv("USERNAME", c.User)
	_ = os.Setenv("USER", c.User)

	c.Home = usr.HomeDir
	_ = os.Setenv("HOME", c.Home)
}
