package cmd

import (
	"fmt"

	"github.com/jgrancell/turtleshell/cli"
)

type Builtins struct {
	Builtins []Builtin
}

type Builtin struct {
	Name        string
	Command     func(*cli.Cli, []string) (int, error)
	Description string
}

func (b *Builtins) Register(n string, c interface{}, d string) {
	builtin := Builtin{
		Name:        n,
		Command:     c.(func(*cli.Cli, []string) (int, error)),
		Description: d,
	}

	b.Builtins = append(b.Builtins, builtin)
}

func LoadBuiltins() *Builtins {
	b := Builtins{}
	b.Register("cd", Changedir, "Changes the shell working directory")
	b.Register("echo", Echo, "Prints the value of a variable")
	b.Register("get", Get, "Gets a shell variable")
	b.Register("history", History, "Provides access to the shell's command history")
	b.Register("home", Home, "Moves to the user's home directory")
	b.Register("redo", Redo, "Runs a previously used command from shell history")
	b.Register("run", Run, "Runs a turtleshell script file")
	b.Register("set", Set, "Sets a shell variable")
	b.Register("status", Status, "Shows the exit status of the last command")
	b.Register("turtle", Turtle, "Displays our friendly mascot, Mr Tutleson")
	return &b
}

func (b *Builtins) Get(name string) (Builtin, error) {
	for _, builtin := range b.Builtins {
		if builtin.Name == name {
			return builtin, nil
		}
	}
	return Builtin{}, fmt.Errorf("builtin with the given name not found")
}
