package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"strings"
)

func part1(replacements map[string][]string, molecule string) {
	unique := make(map[string]struct{})
	for key, list := range replacements {
		n := 0
		for {
			a, b, ok := util.CutAfter(molecule, key, n)
			if !ok {
				break
			}
			for _, value := range list {
				unique[a+value+b] = struct{}{}
			}
			n = len(a) + 1
		}
	}
	fmt.Println("Day 19.1:", len(unique))
}

func part2(replacements map[string][]string, molecule string) {
	steps := 0
	for molecule != "e" {
		for key, list := range replacements {
			for _, value := range list {
				a, b, ok := util.Cut(molecule, value)
				if ok {
					molecule = a + key + b
					steps += 1
				}
			}
		}
	}

	fmt.Println("Day 19.2:", steps)
}

func main() {
	input := puzzle.ReadString(2015, 19)
	values, molecule, _ := util.Cut(input, "\n\n")
	lines := strings.Split(values, "\n")
	replacements := make(map[string][]string, len(lines))
	for _, line := range lines {
		var key, value string
		fmt.Sscanf(line, "%s => %s", &key, &value)

		list, ok := replacements[key]
		if !ok {
			list = make([]string, 0)
		}
		replacements[key] = append(list, value)
	}

	molecule = strings.Trim(molecule, " \r\n\t")
	part1(replacements, molecule)
	part2(replacements, molecule)

}
