package main

import (
    "encoding/json"
    "os"
    "os/user"
    "strings"
)

type Configuration struct {
    PS1     string
    Prompt  string
    User    string
    HomeDir string
    Version string
}

func loadConfiguration() Configuration {
    configuration := Configuration{}
    configuration = parseUser(configuration)
    configuration.Version = "0.0.0-dev"

    configFile := configuration.HomeDir + "/.turtlerc"

    if _, err := os.Stat(configFile); err == nil {
        file, _ := os.Open(configFile)
        defer file.Close()
        decoder := json.NewDecoder(file)
        err := decoder.Decode(&configuration)

        if err != nil {
            configuration.PS1 = "[{user}@{hostname}]: >"
        }
    } else {
        configuration.PS1 = "[{user}@{hostname}]: >"
    }

    configuration.Prompt = parsePrompt(configuration.PS1)

    return configuration
}

func parsePrompt(ps1 string) string {
    shortcodes := make([]string, 3)
    shortcodes[0] = "{user}"
    shortcodes[1] = "{hostname}"
    shortcodes[2] = "{pwd}"

    var prompt string

    // Populating the {cwd} shortcode
    if strings.Contains(ps1, "{cwd}") {
        cwd, _ := os.Getwd()
        prompt = strings.ReplaceAll(ps1, "{cwd}", cwd)
    }

    // Populating the {user} shortcode
    if strings.Contains(prompt, "{user}") {
        user, _ := user.Current()
        prompt = strings.ReplaceAll(prompt, "{user}", user.Username)
        prompt = strings.ReplaceAll(prompt, user.HomeDir, "~")
    }

    // Populating the {hostname} shortcode
    if strings.Contains(prompt, "{hostname}") {
        hostname, _ := os.Hostname()
        prompt = strings.ReplaceAll(prompt, "{hostname}", hostname)
    }

    // Populating all available colors
    // Green
    if strings.Contains(prompt, "{color:red") {
        prompt = strings.ReplaceAll(prompt, "{color:red}", "\033[31m")
        prompt = strings.ReplaceAll(prompt, "{color:red:bold}", "\033[31;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:green") {
        prompt = strings.ReplaceAll(prompt, "{color:green}", "\033[32m")
        prompt = strings.ReplaceAll(prompt, "{color:green:bold}", "\033[32;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:yellow") {
        prompt = strings.ReplaceAll(prompt, "{color:yellow}", "\033[33m")
        prompt = strings.ReplaceAll(prompt, "{color:yellow:bold}", "\033[33;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:blue") {
        prompt = strings.ReplaceAll(prompt, "{color:blue}", "\033[34m")
        prompt = strings.ReplaceAll(prompt, "{color:blue:bold}", "\033[34;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:magenta") {
        prompt = strings.ReplaceAll(prompt, "{color:magenta}", "\033[35m")
        prompt = strings.ReplaceAll(prompt, "{color:magenta:bold}", "\033[35;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:cyan") {
        prompt = strings.ReplaceAll(prompt, "{color:cyan}", "\033[36m")
        prompt = strings.ReplaceAll(prompt, "{color:cyan:bold}", "\033[36;1m")
    }

    // Green
    if strings.Contains(prompt, "{color:white") {
        prompt = strings.ReplaceAll(prompt, "{color:white}", "\033[37m")
        prompt = strings.ReplaceAll(prompt, "{color:white:bold}", "\033[37;1m")
    }

    // Color Reset
    if strings.Contains(prompt, "{endcolor}") {
        prompt = strings.ReplaceAll(prompt, "{endcolor}", "\033[0m")
    }

    return prompt
}

func parseUser(config Configuration) Configuration {
    usr, _ := user.Current()
    config.User = usr.Username
    config.HomeDir = usr.HomeDir
    return config
}
