package main

import (
	"aoc-2015/internal/puzzle"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func isAdventCoin(key string, digit int, mask byte) bool {
	hasher := md5.New()
	hasher.Write([]byte(key))
	hasher.Write([]byte(strconv.Itoa(digit)))
	result := hasher.Sum(nil)
	return result[0] == 0 && result[1] == 0 && (result[2]&mask) == 0
}

func mineCoins(key string, start, count int, mask byte) (int, bool) {
	end := start + count
	for i := start; i < end; i++ {
		if isAdventCoin(key, i, mask) {
			return i, true
		}
	}
	return 0, false
}

func part1(input string) {
	answer := 0
	i := 1
	for {
		if id, ok := mineCoins(input, i, 10000, 0xF0); ok {
			answer = id
			break
		}
		i += 10000
	}
	fmt.Println("Day 04.1:", answer)
}

func part2(input string) {
	i := 1
	n := 1500000
	answer := 0
	for {
		if id, ok := mineCoins(input, i, n, 0xFF); ok {
			answer = id
			break
		}
		i += n
	}
	fmt.Println("Day 04.2:", answer)
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 4), " \r\n\t")
	part1(input)
	part2(input)
}
