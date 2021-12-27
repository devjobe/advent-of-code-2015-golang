package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"strings"
)

type stats struct {
	c, d, a int
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 21), " \r\n\t")
	var hp, dmg, ac int
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d\nArmor: %d", &hp, &dmg, &ac)
	weapons := []stats{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
	armors := []stats{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}
	rings := []stats{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	part1 := 1000
	part2 := 0
	for _, w := range weapons {
		for _, a := range armors {
			for index, r1 := range rings {
				for _, r2 := range rings[index+1:] {
					pdmg := util.MaxInt(1, w.d+a.d+r1.d+r2.d-ac)
					bdmg := util.MaxInt(1, dmg-a.a-r1.a-r2.a)
					cost := w.c + a.c + r1.c + r2.c
					bossRounds := util.MinInt(100, (100+bdmg-1)/bdmg)
					playerRounds := util.MinInt(hp, (hp+pdmg-1)/pdmg)

					if bossRounds >= playerRounds {
						if part1 > cost {
							part1 = cost
						}
					} else if part2 < cost {
						part2 = cost
					}
				}
			}
		}
	}

	fmt.Println("Day 21.1:", part1)
	fmt.Println("Day 21.2:", part2)
}
