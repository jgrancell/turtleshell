package utils

import "strings"

func UnsetFlags(components []string) []string {
	newArgs := make([]string, 0)
	for _, arg := range components {
		if strings.Index(arg, "-") != 0 {
			newArgs = append(newArgs, arg)
		}
	}
	return newArgs
}
