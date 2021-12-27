package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"strconv"
	"strings"
)

func run(input []string, a int) int {
	var ip, b int
	for ip >= 0 && ip < len(input) {
		name, args, _ := util.Cut(input[ip], " ")
		a1, a2, _ := util.Cut(args, ", ")

		a1 = strings.TrimLeft(a1, "+")
		a2 = strings.TrimLeft(a2, "+")

		var r *int
		var offset int
		if a1 == "a" {
			r = &a
			if o, err := strconv.ParseInt(a2, 10, 32); err == nil {
				offset = int(o)
			}
		} else if a1 == "b" {
			r = &b
			if o, err := strconv.ParseInt(a2, 10, 32); err == nil {
				offset = int(o)
			}
		} else {
			if o, err := strconv.ParseInt(a1, 10, 32); err == nil {
				offset = int(o)
			}
		}

		switch name {
		case "hlf":
			*r = *r / 2
			ip += 1
		case "tpl":
			*r = *r * 3
			ip += 1
		case "inc":
			*r = *r + 1
			ip += 1
		case "jmp":
			if offset == 0 {
				break
			}
			ip += offset
		case "jie":
			if (*r & 1) == 0 {
				if offset == 0 {
					break
				}
				ip += offset
			} else {
				ip += 1
			}
		case "jio":
			if *r == 1 {
				if offset == 0 {
					break
				}
				ip += offset
			} else {
				ip += 1
			}
		}
	}
	return b
}

func main() {
	input := puzzle.ReadLines(2015, 23)
	fmt.Println("Day 23.1:", run(input, 0))
	fmt.Println("Day 23.2:", run(input, 1))
}
