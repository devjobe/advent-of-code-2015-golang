package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"sort"
)

type inst struct {
	cmd            string
	x0, y0, x1, y1 int
}

func part1(cmds []inst, X []int, Y []int) {

	grid := make([][]bool, len(X))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, len(grid))
	}

	for _, cmd := range cmds {
		x0 := sort.SearchInts(X, cmd.x0)
		x1 := sort.SearchInts(X, cmd.x1)
		y0 := sort.SearchInts(Y, cmd.y0)
		y1 := sort.SearchInts(Y, cmd.y1)

		if cmd.cmd == "on" {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					grid[x][y] = true
				}
			}
		} else if cmd.cmd == "off" {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					grid[x][y] = false
				}
			}
		} else {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	sum := int64(0)
	for x, row := range grid[:len(grid)-1] {
		for y, v := range row[:len(row)-1] {
			if v {
				sum += int64((X[x+1] - X[x]) * (Y[y+1] - Y[y]))
			}
		}
	}

	fmt.Println("Day 06.1:", sum)
}

func part2(cmds []inst, X []int, Y []int) {
	grid := make([][]int, len(X))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(grid))
	}

	for _, cmd := range cmds {
		x0 := sort.SearchInts(X, cmd.x0)
		x1 := sort.SearchInts(X, cmd.x1)
		y0 := sort.SearchInts(Y, cmd.y0)
		y1 := sort.SearchInts(Y, cmd.y1)

		if cmd.cmd == "on" {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					grid[x][y] += 1
				}
			}
		} else if cmd.cmd == "off" {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					if grid[x][y] > 0 {
						grid[x][y] = grid[x][y] - 1
					}
				}
			}
		} else {
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					grid[x][y] += 2
				}
			}
		}
	}

	sum := int64(0)
	for x, row := range grid[:len(grid)-1] {
		for y, v := range row[:len(row)-1] {
			if v > 0 {
				sum += int64(v) * int64((X[x+1]-X[x])*(Y[y+1]-Y[y]))
			}
		}
	}

	fmt.Println("Day 06.2:", sum)
}

func main() {
	input := puzzle.ReadLines(2015, 6)

	cmds := make([]inst, len(input))
	X := make([]int, len(input)*2)
	Y := make([]int, len(input)*2)

	for index, line := range input {
		var cmd0, cmd1 string
		var x0, x1, y0, y1 int
		n, _ := fmt.Sscanf(line, "%s %d,%d through %d,%d", &cmd1, &x0, &y0, &x1, &y1)
		if n != 5 {
			fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &cmd0, &cmd1, &x0, &y0, &x1, &y1)
		}
		x1 += 1
		y1 += 1
		cmds[index] = inst{cmd1, x0, y0, x1, y1}
		X[index*2] = x0
		X[index*2+1] = x1
		Y[index*2] = y0
		Y[index*2+1] = y1
	}

	sort.Ints(X)
	sort.Ints(Y)

	part1(cmds, X, Y)
	part2(cmds, X, Y)
}
