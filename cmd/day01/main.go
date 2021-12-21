package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"strings"
)

func part1(input string) {
	opening := strings.Count(input, "(")
	answer := 2*opening - len(input)
	fmt.Println("Day 01.1:", answer)
}

func part2(input string) {
	var level = 0
	var answer = 0
	for index, letter := range input {
		if letter == '(' {
			level += 1
		} else if level == 0 {
			answer = index + 1
			break
		} else {
			level -= 1
		}
	}
	fmt.Println("Day 01.2:", answer)
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 1), " \r\n\t")
	part1(input)
	part2(input)
}
