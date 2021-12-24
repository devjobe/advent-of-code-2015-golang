package main

import (
	"aoc-2015/internal/puzzle"
	"fmt"
)

func findRoutes(routes map[string]map[string]int,
	unvisited map[string]struct{},
	current string, current_cost int,
	shortest *int, longest *int) {

	if len(unvisited) == 0 {

		if *shortest == 0 || *shortest > current_cost {
			*shortest = current_cost
		}
		if *longest < current_cost {
			*longest = current_cost
		}
		return
	}

	destinations := routes[current]

	for city, cost := range destinations {
		if _, ok := unvisited[city]; ok {
			delete(unvisited, city)
			findRoutes(routes, unvisited, city, current_cost+cost, shortest, longest)
			unvisited[city] = struct{}{}
		}
	}
}

func main() {
	input := puzzle.ReadLines(2015, 9)

	routes := make(map[string]map[string]int)
	all := make(map[string]struct{})
	for _, line := range input {
		var from, to string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)

		{
			destinations, ok := routes[from]
			if !ok {
				destinations = make(map[string]int)
				routes[from] = destinations
			}
			destinations[to] = distance
		}
		{
			destinations, ok := routes[to]
			if !ok {
				destinations = make(map[string]int)
				routes[to] = destinations
			}
			destinations[from] = distance

		}

		all[to] = struct{}{}
		all[from] = struct{}{}
	}

	var shortest, longest int

	for city := range routes {
		delete(all, city)
		findRoutes(routes, all, city, 0, &shortest, &longest)
		all[city] = struct{}{}
	}

	fmt.Println("Day 09.1:", shortest)
	fmt.Println("Day 09.2:", longest)
}
