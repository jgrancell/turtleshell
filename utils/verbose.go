package utils

func HasVerbose(components []string) bool {
	for _, arg := range components {
		if arg == "--verbose" || arg == "-v" {
			return true
		}
	}
	return false
}
