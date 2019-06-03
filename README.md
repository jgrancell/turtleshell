# TurtleShell

A not very good Linux shell written in Go.

## Objectives

One of the best ways I've found to learn a programming language is to create a
passion project in it, and learn the language as you write and iterate the codebase
for that project.

TurtleShell is a coding experiment to create a full Linux shell in Go, as a
way to learn Go and scratch an itch that I've had for some time -- learn how
to write a Linux shell. Two birds, one stone, etc etc.

TurtleShell is nominally based on the functionality (not the design) of BASH,
however I've made changes to the workflow and functionality to fit my own preferences.
Changes include:

* Builtin commands that make more sense in plain English.
    * Example: `gohome` instead of having to do `cd ~`
    * Example: `get exitcode` instead of having to do `echo $?`
    * Example: `set $VARIABLE value` instead of having to do `VARIABLE=value`
* Easier and more readable customization of shell prompts.
    * Example: `<{color:green}{user}{endcolor}@{cwd}> ` instead of `\[\e[0;32m\u\e[0;0m@\h\]`

## Current and Planned Features

This is still a very, **very** rough development application. I would hesitate to
even begin to label this as pre-alpha. That being said, the features that exist
all work well.

Currently implemented features:
 * The ability to invoke normal system commands and binaries.
 * The mechanisms to create and use shell builtin commands.
     * A number of commands already work, including:
         * `cd` -- Change shell working directory.
         * `exitcodes` -- Show all exit codes currently used by the shell, and their meanings.
         * `get` -- Display a specific environmental variable.
         * `gohome` -- Change the shell working directory to the user's Home Directory.
         * `set` -- Sets a shell variable.
    * The mechanism to create shell builtins is well defined, and easily iterate-able.
* The ability to read arbitrary configuration from `~/.turtlerc` files.
* The ability to customize the turtleshell prompt, which updates in real time while using the shell.
* History saving, including optional timestamps.

Planned features include, but aren't necessarily limited to:

* Setting environment variables via `set VARIABLE value` and `set VARIABLE "multi word value"`.
* Command history with select-and-execute functionality.

## Contributing Suggestions, Issues, and Code

If you are interested in submitting a Feature Request or Enhancement, please submit a
[Github issue](https://github.com/jgrancell/turtleshell/issues). I'm happy to entertain
all Feature Requests and Enhancement suggestions.

If you find a bug, please submit a [Github issue](https://github.com/jgrancell/turtleshell/issues).
If you believe that you've found a security flaw:
1. You're probably right.
2. There's no need to privately message me, because literally no one should be using this code
for anything meaningful.

I will **not** be accepting merge requests to fix bugs or implement features or
enhancements. As said at the beginning of this README.md, this project is a way
for me to learn Go while writing my own shell -- accepting code from others would
defeat that purpose entirely. **I'll happily read over your code to improve my
understanding and credit you accordingly, but the Merge Request itself will be denied.**

## Using TurtleShell

You shouldn't. You really really shouldn't. It's most likely not secure, it's
definitely poorly written in its current state, and i'm prone to rewrite large
portions of the codebase as I increase my knowledge and understanding of Go.

**Do not use this on any system that you care about. Period.**

If you would like to try out TurtleShell, you can in one of two ways:

1. Run the Docker container. `docker run -it jgrancell/turtleshell:development /bin/tsh`. Have fun. Enjoy the Ninja Turtles' default theme.
2. Clone down this repository, move into the `src/` directory, and run `go build -o turtleshell`.
    1. Run TurtleShell via `./turtleshell`, which will run it as a subshell of your current shell.
    2. Alternatively, create a new system user, set that user's default shell to the
    full path of wherever you put the turtleshell binary, and then su/sudo/login to
    that user.

Congratulations, you're now one of the handful of people (a very mostly empty handful)
that have ever used TurtleShell.
