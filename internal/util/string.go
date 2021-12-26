package util

import "strings"

func Cut(s, sep string) (string, string, bool) {
	if index := strings.Index(s, sep); index >= 0 {
		return s[:index], s[index+len(sep):], true
	}

	return s, "", false
}
