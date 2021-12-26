package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"math"
	"strings"
)

func findOptimalArrangement(happiness map[string]map[string]int, first, guest string,
	seated map[string]struct{}, current int, optimal *int) {

	if len(happiness) == len(seated) {
		p1 := happiness[guest][first]
		p2 := happiness[first][guest]
		result := current + p1 + p2
		if result > *optimal {
			*optimal = result
		}
		return
	}

	list, ok := happiness[guest]
	if !ok {
		return
	}

	for name, points := range list {
		if _, ok := seated[name]; !ok {
			seated[name] = struct{}{}
			points2 := happiness[name][guest]
			findOptimalArrangement(happiness, name, first, seated, current+points+points2, optimal)
			delete(seated, name)
		}
	}
}

func calculateOptimal(happiness map[string]map[string]int) int {
	optimal := math.MinInt
	seated := make(map[string]struct{}, len(happiness))
	for name := range happiness {
		seated[name] = struct{}{}
		findOptimalArrangement(happiness, name, name, seated, 0, &optimal)
		delete(seated, name)
	}
	return optimal
}

func main() {
	input := puzzle.ReadLines(2015, 13)

	all := make(map[string]int)
	happiness := make(map[string]map[string]int, len(input))
	for _, line := range input {
		var n1, n2, verb string
		var points int
		fmt.Sscanf(strings.TrimRight(line, "."), "%s would %s %d happiness units by sitting next to %s", &n1, &verb, &points, &n2)
		if verb == "lose" {
			points = -points
		}

		names, ok := happiness[n1]
		if !ok {
			names = make(map[string]int)
			happiness[n1] = names
		}
		names[n2] = points
		all[n1] = 0
		all[n2] = 0
	}

	fmt.Println("Day 13.1:", calculateOptimal(happiness))
	for name := range all {
		happiness[name]["ME"] = 0
	}
	happiness["ME"] = all

	fmt.Println("Day 13.2:", calculateOptimal(happiness))
}
