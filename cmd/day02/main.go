package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
)

func requiredPaper(box []int64) int64 {
	l, w, h := box[0], box[1], box[2]
	a1 := l * w
	a2 := l * h
	a3 := w * h
	return util.Min(a1, a2, a3) + 2*(a1+a2+a3)
}

func requiredRibbon(box []int64) int64 {
	l, w, h := box[0], box[1], box[2]
	s1 := 2 * (l + w)
	s2 := 2 * (l + h)
	s3 := 2 * (w + h)
	return util.Min(s1, s2, s3) + l*w*h
}

func part1(input [][]int64) {
	var answer int64
	for _, box := range input {
		answer += requiredPaper(box)
	}
	fmt.Println("Day 02.1:", answer)
}

func part2(input [][]int64) {
	var answer int64
	for _, box := range input {
		answer += requiredRibbon(box)
	}
	fmt.Println("Day 02.2:", answer)
}

func main() {
	input := puzzle.IntegerMatrix(puzzle.ReadString(2015, 2), "x")
	part1(input)
	part2(input)
}
