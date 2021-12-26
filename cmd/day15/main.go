package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
)

func getScore(properties [][5]int, value []int) int {
	score := 1
	for i := 0; i < 4; i++ {
		val := 0
		for k := 0; k < len(properties); k++ {
			val += properties[k][i] * value[k]
		}
		if val <= 0 {
			score = 0
			break
		}
		score *= val
	}
	return score
}

func findScore(properties [][5]int, value []int, calories int) int {
	sum := 0
	for _, n := range value {
		sum += n
	}
	n := 100 - sum
	k := len(value)
	value = append(value, n)

	score := 0
	if k == 3 {
		if calories > 0 {
			val := 0
			for p := 0; p < len(properties); p++ {
				val += properties[p][4] * value[p]
			}
			if val != calories {
				return 0
			}
		}
		return getScore(properties, value)
	}

	n -= 3 - k
	for i := 1; i <= n; i++ {
		value[k] = i
		s := findScore(properties, value, calories)
		if s > score {
			score = s
		}
	}

	return score
}

func main() {
	input := puzzle.ReadLines(2015, 15)

	properties := make([][5]int, len(input))

	for index, line := range input {
		var d0, d1, d2, d3, d4 int
		_, props, _ := util.Cut(line, ": ")
		fmt.Sscanf(props, "capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&d0, &d1, &d2, &d3, &d4)

		properties[index] = [5]int{d0, d1, d2, d3, d4}
	}

	value := [4]int{0, 0, 0, 0}
	score := findScore(properties, value[0:0], 0)
	fmt.Println("Day 15.1:", score)
	score2 := findScore(properties, value[0:0], 500)
	fmt.Println("Day 15.2:", score2)
}
