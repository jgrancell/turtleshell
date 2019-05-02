package main

type Builtins struct {
    Commands  map[string]func([]string, Configuration)
    Help      map[string]string
}

func NewBuiltins() Builtins {
    commands := map[string]func([]string, Configuration) {
        "cd": builtins_cd,
        "exitcodes": builtins_exitcodes,
        "get": builtins_get,
        "gohome": builtins_gohome,
        "turtle": builtins_turtle,
    }

    help := map[string]string {
        "cd": "Changes the shell working directory.",
        "exitcodes": "Show all shell exitcodes.",
        "get": "Gets an environmental variable.",
        "gohome": "Move to the user's home directory.",
    }

    b := Builtins {commands, help}
    return b
}

func isBuiltin(command string) bool {
    builtins := NewBuiltins()
    for k, _ := range builtins.Commands {
        if command == k {
            return true
        }
    }
    return false
}