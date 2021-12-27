package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func divisiorSum(n int64) int64 {
	if n == 1 {
		return 1
	}
	sum := n + 1
	max := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= max; i++ {
		if (n % i) == 0 {
			k := n / i
			if i == k {
				sum += i
			} else {
				sum += i + k
			}
		}
	}
	return sum
}

func part2Sum(n int64) int64 {
	if n < 51 {
		return divisiorSum(n)
	}
	sum := n
	max := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= max; i++ {
		if (n % i) == 0 {
			if i*50 >= n {
				sum += i
			}
			k := n / i
			if k != i && k*50 >= n {
				sum += k
			}
		}
	}
	return sum
}

func main() {
	input, _ := strconv.ParseInt(strings.Trim(puzzle.ReadString(2015, 20), " \r\n\t"), 10, 64)

	n := int64(1)
	for divisiorSum(n)*10 < input {
		n += 1
	}
	fmt.Println("Day 20.1:", n)

	n = int64(1)
	for part2Sum(n)*11 < input {
		n += 1
	}
	fmt.Println("Day 20.2:", n)
}
