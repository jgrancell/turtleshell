package main

func runnerIdentify(args []string) map[string]string {
	array := make(map[string]string)

	if args[0] == "run" || args[0] == "." {
		array["runner"] = "file"
		array["file"] = args[1]
	} else if args[0] == "-c" || args[0] == "--command" {
		array["runner"] = "command"
		array["command"] = args[1]
	} else {
		array["runner"] = "binary"
	}

	return array
}
