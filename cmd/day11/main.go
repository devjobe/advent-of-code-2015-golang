package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
	"strings"
)

func incrementPassword(password []byte) {
	const min = byte('a')
	const max = byte('z')
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] == max {
			password[i] = min
		} else {
			password[i] = password[i] + 1
			break
		}
	}

	for i, ch := range password {
		if ch == 'i' || ch == 'o' || ch == 'l' {
			password[i] = ch + 1
			for j := range password[i+1:] {
				password[j] = min
			}
			break
		}
	}
}

func indexOfSequence(input []byte) int {
	for i, ch := range input[:len(input)-2] {
		if input[i+1] == ch+1 && input[i+2] == ch+2 {
			return i
		}
	}
	return -1
}

func indexOfPair(input []byte) int {
	for i, ch := range input[:len(input)-1] {
		if input[i+1] == ch {
			return i
		}
	}
	return -1
}

func nextPassword(input string) string {
	var password [8]byte
	copy(password[:], input)

	for {
		incrementPassword(password[:])
		sequence := indexOfSequence(password[:])
		pair := indexOfPair(password[:])
		if sequence >= 0 && pair >= 0 && indexOfPair(password[pair+2:]) >= 0 {
			return string(password[:])
		}

		if (pair < 0 || pair >= 3) && (sequence < 0 || sequence >= 4) {
			n := password[3]
			if n > byte('x') ||
				password[4] >= n ||
				password[5] >= n+1 ||
				password[6] >= n+2 ||
				password[7] >= n+2 {
				if n >= byte('x') {
					password[3] = 'z'
				}
				incrementPassword(password[:4])
				n = password[3]
			}
			password[4] = n
			password[5] = n + 1
			password[6] = n + 2
			password[7] = n + 2
			return string(password[:])
		}
	}
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 11), " \r\n\t")
	part1 := nextPassword(input)
	fmt.Println("Day 11.1:", part1)
	part2 := nextPassword(part1)
	fmt.Println("Day 11.2:", part2)
}
