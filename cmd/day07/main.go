package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"log"
	"strconv"
)

func main() {
	input := puzzle.ReadLines(2015, 7)
	output := make(map[string]uint16)

	has_b := false
	b := uint16(0)

	value := func(s string) (uint16, bool) {
		d, err := strconv.Atoi(s)
		if err == nil {
			return uint16(d), true
		}

		if has_b && s == "b" {
			return b, true
		}

		v, ok := output[s]
		return v, ok
	}

	pending := make([]string, 0, len(input))
	remaining := input
	for {
		if len(remaining) == 0 {
			if len(pending) == 0 {
				if !has_b {
					has_b = true
					b = output["a"]
					output = make(map[string]uint16)
					pending = input
				} else {
					break
				}
			}
			remaining, pending = pending, make([]string, 0, len(input))
		}

		var line string
		line, remaining = remaining[0], remaining[1:]

		var a1, cmd, a2, res string
		n, _ := fmt.Sscanf(line, "%s %s %s -> %s", &a1, &cmd, &a2, &res)
		if n != 4 {
			n, _ := fmt.Sscanf(line, "%s %s -> %s", &cmd, &a1, &res)
			if n != 3 {
				fmt.Sscanf(line, "%s -> %s", &a1, &res)
				if v, ok := value(a1); ok {
					output[res] = v
				} else {
					pending = append(pending, line)
				}
			} else if cmd == "NOT" {
				if v, ok := value(a1); ok {
					output[res] = ^v
				} else {
					pending = append(pending, line)
				}
			} else {
				log.Fatal("Unexpected input:", line)
			}
		} else {

			v1, ok1 := value(a1)
			v2, ok2 := value(a2)
			if !ok1 || !ok2 {
				pending = append(pending, line)
			} else if cmd == "AND" {
				output[res] = v1 & v2
			} else if cmd == "OR" {
				output[res] = v1 | v2
			} else if cmd == "LSHIFT" {
				output[res] = v1 << v2
			} else if cmd == "RSHIFT" {
				output[res] = v1 >> v2
			} else {
				log.Fatal("Unexpected input:", line)
			}
		}
	}

	fmt.Println("Day 07.1:", b)
	fmt.Println("Day 07.2:", output["a"])
}
