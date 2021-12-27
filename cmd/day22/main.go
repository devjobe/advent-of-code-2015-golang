package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
	"math"
	"strings"
)

type Stats struct {
	spent                    int
	hp, mana, armor          int
	shield, poison, recharge int
	bossHp, bossDmg          int
	hard                     bool
}

const (
	MAGIC_MISSILE = iota
	DRAIN
	SHIELD
	POISON
	RECHARGE
	SPELL_COUNT
)

func (s *Stats) castSpell(n int) bool {
	spellCost := [6]int{53, 73, 113, 173, 229}[n]
	if s.mana < spellCost {
		return false
	}
	s.spent += spellCost
	s.mana -= spellCost
	switch n {
	case MAGIC_MISSILE:
		s.bossHp -= 4
	case DRAIN:
		s.bossHp -= 2
		s.hp += 2
	case SHIELD:
		if s.shield > 0 {
			return false
		}

		s.shield = 6
	case POISON:
		if s.poison > 0 {
			return false
		}
		s.poison = 6
	case RECHARGE:
		if s.recharge > 0 {
			return false
		}
		s.recharge = 5
	}

	return true
}

func (s *Stats) tick() {
	if s.poison > 0 {
		s.poison -= 1
		s.bossHp -= 3
	}

	if s.shield > 0 {
		s.shield -= 1
		s.armor = 7
	} else {
		s.armor = 0
	}

	if s.recharge > 0 {
		s.recharge -= 1
		s.mana += 101
	}
}

func (s *Stats) bossTurn() {
	s.hp -= util.MaxInt(1, s.bossDmg-s.armor)
}

func (s *Stats) round(best *int) {
	if s.hard {
		s.hp -= 1
		if s.hp <= 0 {
			return
		}
	}
	for i := 0; i < 5; i++ {
		stats := *s
		stats.tick()
		if stats.castSpell(i) && stats.spent < *best {
			stats.tick()
			if stats.bossHp > 0 {
				stats.bossTurn()
				if stats.hp > 0 {
					stats.round(best)
				}
			} else if *best > stats.spent {
				*best = stats.spent
			}
		}
	}
}

func main() {
	input := strings.Trim(puzzle.ReadString(2015, 22), " \r\n\t")
	var hp, dmg int
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d", &hp, &dmg)

	var stats Stats
	stats.hp = 50
	stats.mana = 500
	stats.bossHp = hp
	stats.bossDmg = dmg

	var best int = math.MaxInt
	stats.round(&best)
	if best != math.MaxInt {
		fmt.Println("Day 22.1:", best)
	} else {
		fmt.Println("Failed to find a solution")
	}

	best = math.MaxInt
	stats.hard = true
	stats.round(&best)
	if best != math.MaxInt {
		fmt.Println("Day 22.2:", best)
	} else {
		fmt.Println("Failed to find a solution")
	}
}
