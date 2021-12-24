package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"strings"
)

func lookAndSay(s string) string {
	var count int
	var prev rune

	builder := strings.Builder{}

	for _, ch := range s {
		if ch == prev {
			count += 1
		} else {
			if prev != rune(0) {
				builder.WriteRune(rune(count + int('0')))
				builder.WriteRune(prev)
			}
			prev = ch
			count = 1
		}
	}
	builder.WriteRune('0' + rune(count))
	builder.WriteRune(prev)

	return builder.String()
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 10), " \n\r\t")

	s := input
	for i := 0; i < 40; i++ {
		s = lookAndSay(s)
	}
	fmt.Println("Day 10.1:", len(s))

	for i := 0; i < 10; i++ {
		s = lookAndSay(s)
	}
	fmt.Println("Day 10.2:", len(s))
}
