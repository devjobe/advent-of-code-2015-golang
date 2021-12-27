package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
)

func getNumber(n int64) int64 {
	var value int64 = 20151125
	const (
		factor  = 252533
		divisor = 33554393
	)
	for ; n > 1; n-- {
		value = (value * factor) % divisor
	}
	return value
}

func triangleSum(n int64) int64 {
	return n * (n + 1) / 2
}

func main() {
	input := puzzle.ReadInt64List(2015, 25)
	if len(input) != 2 {
		fmt.Println("Expected two numbers")
		return
	}

	r, c := input[0], input[1]
	fmt.Println("Day 25:", getNumber(triangleSum(c+r-1)-r+1))
}
