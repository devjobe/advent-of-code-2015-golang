package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"math"
	"strings"
)

func findSmallest(list []int64, weight, product int64, best *int64) {
	for index, n := range list {
		if n > weight {
			continue
		}
		if math.MaxInt64/n <= product {
			continue
		}
		q := product * n
		if q >= *best {
			continue
		}
		if n == weight {
			*best = q
		} else {
			findSmallest(list[index+1:], weight-n, q, best)
		}
	}
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 24), " \r\n\t")
	list := puzzle.IntegerList(input, "\n")
	{
		weight := util.Sum(list...) / 3
		var best int64 = math.MaxInt64
		findSmallest(list, weight, 1, &best)
		fmt.Println("Day 24.1:", best)
	}
	{
		weight := util.Sum(list...) / 4
		var best int64 = math.MaxInt64
		findSmallest(list, weight, 1, &best)
		fmt.Println("Day 24.2:", best)
	}
}
