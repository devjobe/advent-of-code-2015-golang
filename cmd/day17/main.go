package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"strconv"
)

func getContainers(capacity int, containers []int, depth int, limitDepth int) (int, int) {
	if limitDepth > 0 && depth > limitDepth {
		return 0, 0
	}
	if capacity == 150 {
		return 1, depth
	} else if capacity > 150 || len(containers) == 0 {
		return 0, 0
	}

	count := 0
	minDepth := 0
	for i := 0; i < len(containers); i++ {
		n, d := getContainers(capacity+containers[i], containers[i+1:], depth+1, limitDepth)
		count += n

		if minDepth == 0 || (d > 0 && minDepth > d) {
			minDepth = d
		}
	}

	return count, minDepth
}

func main() {
	input := puzzle.ReadLines(2015, 17)
	values := make([]int, len(input))
	for index, line := range input {
		v, _ := strconv.Atoi(line)
		values[index] = v
	}

	count, minDepth := getContainers(0, values, 0, 0)
	fmt.Println("Day 17.1:", count)
	count2, _ := getContainers(0, values, 0, minDepth)
	fmt.Println("Day 17.2:", count2)
}
