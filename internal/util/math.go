package util

func Min(array ...int64) int64 {
	min := array[0]
	for _, n := range array[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
