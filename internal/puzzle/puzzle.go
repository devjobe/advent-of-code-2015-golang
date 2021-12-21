package puzzle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

func ReadString(year, day int) string {
	return string(fetchOrReadInput(year, day))
}
