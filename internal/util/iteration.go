package util

func AnyWindow(input string, window_size int, predicate func(string) bool) bool {
	if len(input) < window_size {
		return false
	}
	end := len(input) - window_size
	for i := 0; i <= end; i++ {
		if predicate(input[i : i+window_size]) {
			return true
		}
	}
	return false
}
