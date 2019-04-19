package main

type Builtins struct {
    Commands  map[string]func([]string, Configuration)
    Help      map[string]string
}

func NewBuiltins() Builtins {
    commands := map[string]func([]string, Configuration) {
        "cd": cd,
        "gohome": gohome,
        "get": get,
        "exitcodes": exitcodes,
    }

    help := map[string]string {
        "cd": "Changes the shell working directory.",
        "gohome": "Move to the user's home directory.",
        "get": "Gets an environmental variable.",
        "exitcodes": "Show all shell exitcodes.",
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
