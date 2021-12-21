package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"strings"
)

type Coords struct {
	x, y int
}

func part1(input string) {
	houses := make(map[Coords]struct{}, len(input))
	current := Coords{x: 0, y: 0}
	visited := struct{}{}
	houses[current] = visited

	for _, ch := range input {
		switch ch {
		case '^':
			current.y += 1
		case 'v':
			current.y -= 1
		case '>':
			current.x += 1
		case '<':
			current.x -= 1
		}
		houses[current] = visited
	}
	answer := len(houses)
	fmt.Println("Day 03.1:", answer)
}

func part2(input string) {
	houses := make(map[Coords]struct{}, len(input))
	santa := Coords{x: 0, y: 0}
	robosanta := Coords{x: 0, y: 0}
	visited := struct{}{}
	houses[santa] = visited

	coords := [2]*Coords{&santa, &robosanta}

	for index, ch := range input {
		current := coords[index&1]
		switch ch {
		case '^':
			current.y += 1
		case 'v':
			current.y -= 1
		case '>':
			current.x += 1
		case '<':
			current.x -= 1
		}

		houses[*current] = visited
	}
	answer := len(houses)
	fmt.Println("Day 03.2:", answer)
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 3), " \r\n\t")
	part1(input)
	part2(input)
}
