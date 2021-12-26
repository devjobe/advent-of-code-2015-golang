package puzzle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc-2015/internal/session"
)

func fetchAndSave(year, day int, path string) ([]byte, error) {
	input, err := session.FetchInput(year, day)
	if err != nil {
		return nil, err
	}

	ioutil.WriteFile(path, input, 0644)
	return input, nil
}

func fetchOrReadInput(year, day int) []byte {
	path := fmt.Sprintf("inputs/%d-day%02d.txt", year, day)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			content, err = fetchAndSave(year, day, path)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return content
}

func ReadBytes(year, day int) []byte {
	return fetchOrReadInput(year, day)
}

func ReadString(year, day int) string {
	return string(fetchOrReadInput(year, day))
}

func ReadLines(year, day int) []string {
	input := strings.Trim(ReadString(year, day), " \r\n\t")
	return strings.Split(input, "\n")
}

func IntegerList(line, sep string) []int64 {
	input := strings.Trim(line, " \r\n\t")
	split := strings.Split(input, sep)
	result := make([]int64, len(split))
	for index, input := range split {
		i, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			log.Fatal("Error parsing integer list", err)
		}
		result[index] = i
	}
	return result
}

func IntegerMatrix(s string, sep string) [][]int64 {
	input := strings.Trim(s, " \r\n\t")
	lines := strings.Split(input, "\n")
	result := make([][]int64, len(lines))
	for index, line := range lines {
		result[index] = IntegerList(line, sep)
	}
	return result
}
