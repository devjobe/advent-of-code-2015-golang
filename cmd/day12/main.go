package main

import (
	"aoc-2015/internal/puzzle"
	"encoding/json"
	"fmt"
)

func countNumbers(entry interface{}, ignoreRed bool) int64 {
	sum := int64(0)

	switch entry.(type) {
	case []interface{}:
		for _, v := range entry.([]interface{}) {
			if f, ok := v.(float64); ok {
				sum += int64(f)
			} else {
				sum += countNumbers(v, ignoreRed)
			}
		}
	case map[string]interface{}:
		obj := entry.(map[string]interface{})
		for _, v := range obj {
			if f, ok := v.(float64); ok {
				sum += int64(f)
			} else if s, ok := v.(string); ok {
				if ignoreRed && s == "red" {
					return 0
				}
			} else {
				sum += countNumbers(v, ignoreRed)
			}
		}
	}
	return sum
}

func main() {
	input := puzzle.ReadBytes(2015, 12)
	var payload []interface{}
	err := json.Unmarshal(input, &payload)
	if err != nil {
		fmt.Println(err)
	}

	sum := countNumbers(payload, false)
	fmt.Println("Day 12.1:", sum)
	sum = countNumbers(payload, true)
	fmt.Println("Day 12.2:", sum)
}
