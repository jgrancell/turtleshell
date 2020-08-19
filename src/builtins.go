package main

// Builtins maps command line built in commands to go functions.
type Builtins struct {
	Commands map[string]func([]string, Configuration)
	Help     map[string]string
}

// NewBuiltins generates a Builtins struct for application use.
func NewBuiltins() Builtins {
	commands := map[string]func([]string, Configuration){
		"cd":        builtinsCd,
		"echo":      builtinsEcho,
		"exitcodes": builtinsExitcodes,
		"get":       builtinsGet,
		"gohome":    builtinsGohome,
		"history":   builtinsHistory,
		"redo":      builtinsRedo, // part of the builtins_history file
		"set":       builtinsSet,
		"turtle":    builtinsTurtle,
	}

	help := map[string]string{
		"cd":        "Changes the shell working directory.",
		"echo":      "Prints the value of a v",
		"exitcodes": "Show all shell exitcodes.",
		"get":       "Gets an environmental variable.",
		"gohome":    "Move to the user's home directory.",
		"history":   "Shows a list of all shell history items.",
		"redo":      "Runs a previously run command from history.",
		"set":       "Sets an environmental variable.",
	}

	b := Builtins{commands, help}
	return b
}

func isBuiltin(command string) bool {
	builtins := NewBuiltins()
	for k := range builtins.Commands {
		if command == k {
			return true
		}
	}
	return false
}
