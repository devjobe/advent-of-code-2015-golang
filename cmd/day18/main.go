package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
)

func countNeighbours(grid [][]bool, x0, y0 int) int {
	y1 := util.MinInt(y0+1, len(grid)-1)
	n := 0
	for y := util.MaxInt(y0-1, 0); y <= y1; y++ {
		x1 := util.MinInt(x0+1, len(grid[y])-1)
		for x := util.MaxInt(x0-1, 0); x <= x1; x++ {
			if grid[y][x] {
				n += 1
			}

		}
	}
	return n
}

func counts(input []string, sticky bool) int {
	grid := make([][]bool, len(input))
	state := make([][]bool, len(grid))
	for index, line := range input {
		list := make([]bool, len(line))
		for n, ch := range line {
			if ch == '#' {
				list[n] = true
			}
		}
		grid[index] = list
		state[index] = make([]bool, len(list))
	}
	x1 := len(grid[0]) - 1
	y1 := len(grid) - 1
	x2 := len(grid[y1]) - 1

	if sticky {
		grid[0][0] = true
		grid[0][x1] = true
		grid[y1][0] = true
		grid[y1][x2] = true

		state[0][0] = true
		state[0][x1] = true
		state[y1][0] = true
		state[y1][x2] = true
	}

	for step := 1; step <= 100; step++ {
		g := grid
		grid = state
		state = g
		for index, row := range g {
			res := grid[index]
			for k, v := range row {
				if sticky {
					if index == 0 {
						if k == 0 || k == x1 {
							continue
						}
					}
					if index == y1 {
						if k == 0 || k == x2 {
							continue
						}
					}
				}
				count := countNeighbours(g, k, index)
				if v {
					if count == 3 || count == 4 {
						res[k] = true
					} else {
						res[k] = false
					}
				} else if count == 3 {
					res[k] = true
				} else {
					res[k] = false
				}
			}
		}
	}

	on := 0
	for _, row := range grid {
		for _, v := range row {
			if v {
				on += 1
			}
		}
	}
	return on
}

func main() {
	input := puzzle.ReadLines(2015, 18)

	fmt.Println("Day 18.1:", counts(input, false))
	fmt.Println("Day 18.2:", counts(input, true))
}
