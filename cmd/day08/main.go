package main

import (
	"aoc-2015/internal/puzzle"
	"strings"
)

func main() {
	input := puzzle.ReadLines(2015, 8)

	var decoded int
	for _, line := range input {

		decoded += 2

		s := line[1 : len(line)-1]
		for index := strings.Index(s, "\\"); index >= 0; index = strings.Index(s, "\\") {
			if s[index+1] == 'x' {
				index += 4
				decoded += 3
			} else {
				index += 2
				decoded += 1
			}

			if index >= len(s) {
				break
			}
			s = s[index:]
		}

	}

	println("Day 08.1:", decoded)

	var encoded int
	for _, line := range input {
		encoded += 2 + strings.Count(line, "\\") + strings.Count(line, "\"")
	}
	println("Day 08.2:", encoded)
}
