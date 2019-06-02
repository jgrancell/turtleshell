package main

func runner_identify(args []string) map[string]string {
    array := make(map[string]string)

    if args[0] == "run" || args[0] == "." {
        array["runner"] = "file"
    } else {
        array["runner"] = "binary"
    }

    return array
}
