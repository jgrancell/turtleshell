# TurtleShell

A not very good Linux shell written in Go.

[![build](https://img.shields.io/travis/com/jgrancell/turtleshell?style=flat-square&logo=travis)](https://travis-ci.com/github/jgrancell/turtleshell)
[![coverage](https://img.shields.io/codecov/c/github/jgrancell/turtleshell?style=flat-square&token=K1u8dfG8TA)](https://codecov.io/gh/jgrancell/turtleshell)

[![releases](https://img.shields.io/github/v/release/jgrancell/turtleshell?style=flat-square)](https://github.com/jgrancell/turtleshell/releases)
[![GitHub license](https://img.shields.io/github/license/jgrancell/turtleshell?color=green&style=flat-square)](https://github.com/jgrancell/turtleshell/blob/main/LICENSE)

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
    * Example: `home` instead of having to do `cd ~`
    * Example: `status` instead of having to do `echo $?`
* Easier and more readable customization of shell prompts.
    * Example: `<{color:green}{user}{endcolor}@{cwd}> ` instead of `\[\e[0;32m\u\e[0;0m@\h\]`

## Current and Planned Features

Turtleshell is under active development, and could now in 2021

This is still a development application. I would hesitate to
even begin to label this as pre-alpha. That being said, the features that exist
all work well.

Currently implemented features:
 * The ability to invoke normal system commands and binaries.
 * The mechanisms to create and use shell builtin commands.
     * A number of commands already work, including:
         * `cd`        -- Changes the shell working directory.
         * `echo`      -- Prints the value of a variable.
         * `get`       -- Display a specific environmental variable.
         * `home`      -- Change the shell working directory to the user's Home Directory.
         * `history`   -- Shows a list of all items that have been run in this shell session.
         * `info`      -- Shows shell version and diagnostic information
         * `run`       -- Runs arbitrary scripts using shellang, the scripting language that supports turtleshell.
         * `set` -- Sets a shell variable.
         * `status`    -- Show the exit status of the last command.
         * `turtle`    -- Prints a very cool turtle
    * The mechanism to create shell builtins is well defined, and builtins are easy to create.
* The ability to read shell configuration from `~/.turtlerc` files.
* The ability to customize the turtleshell prompt, which updates in real time while using the shell.
* History saving, including optional timestamps.

Planned features include, but aren't necessarily limited to:

* History:
  * ~~Command history with select-and-execute functionality.~~ [execute pending]
  * Being able to hit the up arrow key to walk through history on the command line.
  * `history !` and `history last` just reruns the last command
  * `history !sudo` runs the last command, but with sudo

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
(June 2021 edit: I rewrote the whole thing -- i'll probably do it again, too)

**Do not use this on any system that you care about. Period.**

If you would like to try out TurtleShell, you can in one of two ways:

1. Run the Docker container. `docker run -it jgrancell/turtleshell /bin/tsh`. Have fun. Enjoy the Ninja Turtles' default theme.
2. Clone down this repository and run `make build`.
    1. Run TurtleShell via `./turtleshell`, which will run it as a subshell of your current shell.
    2. Alternatively, create a new system user, set that user's default shell to the
    full path of wherever you put the turtleshell binary, and then su/sudo/login to
    that user.

Congratulations, you're now one of the handful of people (a very mostly empty handful)
that have ever used TurtleShell.
