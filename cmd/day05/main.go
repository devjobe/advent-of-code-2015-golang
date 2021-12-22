package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"strings"
)

func isNice(input string) bool {
	banned := [4]string{"ab", "cd", "pq", "xy"}
	for _, word := range banned {
		if strings.Contains(input, word) {
			return false
		}
	}
	const vowels = "aeiou"
	count := 0
	for _, ch := range input {
		if strings.ContainsRune(vowels, ch) {
			count += 1
		}
	}
	if count < 3 {
		return false
	}

	prev := '\u0000'
	for _, ch := range input {
		if prev == ch {
			return true
		}
		prev = ch
	}
	return false
}

func part1(input []string) {
	answer := 0
	for _, line := range input {
		if isNice(line) {
			answer += 1
		}
	}
	fmt.Println("Day 05.1:", answer)
}

func isNice2(input string) bool {
	if !util.AnyWindow(input, 2, func(window string) bool {
		return strings.Count(input, window) >= 2
	}) {
		return false
	}

	prev2 := '\u0000'
	prev1 := '\u0000'
	for _, ch := range input {
		if ch == prev2 {
			return true
		}
		prev2 = prev1
		prev1 = ch
	}
	return false
}

func part2(input []string) {
	var answer = 0
	for _, line := range input {
		if isNice2(line) {
			answer += 1
		}
	}
	fmt.Println("Day 05.2:", answer)
}

func main() {
	input := puzzle.ReadLines(2015, 5)
	part1(input)
	part2(input)
}
