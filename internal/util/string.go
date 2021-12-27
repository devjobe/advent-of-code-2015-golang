package util

import "strings"

func Cut(s, sep string) (string, string, bool) {
	if index := strings.Index(s, sep); index >= 0 {
		return s[:index], s[index+len(sep):], true
	}
	return s, "", false
}

func CutAfter(s, sep string, n int) (string, string, bool) {
	if n >= len(s) {
		return s, "", false
	}
	if index := strings.Index(s[n:], sep); index >= 0 {
		index += n
		return s[:index], s[index+len(sep):], true
	}
	return s, "", false
}
