package utils

import "strings"

func JoinToString(array []string, sep string) string {
	var s string

	for _, elem := range array {
		s += elem
		s += sep
	}

	if len(array) > 0 {
		s = strings.TrimSuffix(s, sep)
	}

	return s
}
