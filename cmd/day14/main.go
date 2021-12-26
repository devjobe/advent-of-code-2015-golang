package main

import (
	"aoc-2015/internal/puzzle"
	"aoc-2015/internal/util"
	"fmt"
)

type reindeer struct {
	points                int64
	speed, duration, rest int64
	name                  string
}

func (entry *reindeer) fromString(line string) {
	fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
		&entry.name, &entry.speed, &entry.duration, &entry.rest)
}

func (e *reindeer) distance(raceDuration int64) int64 {
	cycle := e.duration + e.rest
	n := raceDuration / cycle
	active := e.duration*n + util.Min(raceDuration%cycle, e.duration)
	return e.speed * active
}

func main() {
	input := puzzle.ReadLines(2015, 14)
	maxDistance := int64(0)
	const totalDuration int64 = 2503

	list := make([]reindeer, len(input))
	for index, line := range input {
		entry := &list[index]
		entry.fromString(line)
		d := entry.distance(totalDuration)
		maxDistance = util.Max(d, maxDistance)
	}

	fmt.Println("Day 14.1:", maxDistance)

	distances := make([]int64, len(list))
	for i := int64(1); i <= totalDuration; i++ {
		leadingDistance := int64(0)
		for index := range list {
			distances[index] = list[index].distance(i)
			leadingDistance = util.Max(leadingDistance, distances[index])
		}

		for index := range list {
			if distances[index] == leadingDistance {
				entry := &list[index]
				entry.points += 1
			}
		}
	}

	maxPoints := int64(0)
	for index := range list {
		maxPoints = util.Max(list[index].points, maxPoints)
	}

	fmt.Println("Day 14.2:", maxPoints)
}
