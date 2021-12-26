package util

func Max(array ...int64) int64 {
	max := array[0]
	for _, n := range array[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func Min(array ...int64) int64 {
	min := array[0]
	for _, n := range array[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
