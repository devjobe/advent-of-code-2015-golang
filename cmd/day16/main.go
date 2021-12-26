package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"strings"
)

func main() {
	input := puzzle.ReadLines(2015, 16)

	expected := make(map[string]int)
	expected["children"] = 3
	expected["cats"] = 7
	expected["samoyeds"] = 2
	expected["pomeranians"] = 3
	expected["akitas"] = 0
	expected["vizslas"] = 0
	expected["goldfish"] = 5
	expected["trees"] = 3
	expected["cars"] = 2
	expected["perfumes"] = 1

	match := func(k string, v int) bool {
		if k == "cats" || k == "trees" {
			return expected[k] < v
		} else if k == "pomeranians" || k == "goldfish" {
			return expected[k] > v
		} else {
			return expected[k] == v
		}
	}

	var part1, part2 int
	for _, line := range input {
		var sue, v0, v1, v2 int
		var n0, n1, n2 string

		a, b, _ := util.Cut(line, ": ")
		fmt.Sscanf(a, "Sue %d", &sue)
		fmt.Sscanf(b, "%s %d, %s %d, %s %d", &n0, &v0, &n1, &v1, &n2, &v2)

		n0 = strings.TrimRight(n0, ":")
		n1 = strings.TrimRight(n1, ":")
		n2 = strings.TrimRight(n2, ":")
		if v, ok := expected[n0]; !ok || v == v0 {
			if v, ok := expected[n1]; !ok || v == v1 {
				if v, ok := expected[n2]; !ok || v == v2 {
					part1 = sue
				}
			}
		}

		if match(n0, v0) && match(n1, v1) && match(n2, v2) {
			part2 = sue
		}
	}

	fmt.Println("Day 16.1:", part1)
	fmt.Println("Day 16.2:", part2)
}
