package session

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ReadCookie() (string, string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	content, err := ioutil.ReadFile(filepath.Join(home_dir, ".aoc-session"))
	if err != nil {
		return "", "", err
	}

	split := strings.SplitN(string(content), "=", 2)
	if len(split) != 2 {
		return "", "", errors.New("unknown aoc cookie format")
	}
	return split[0], split[1], nil
}

func FetchInput(year int, day int) ([]byte, error) {
	session_name, session_value, err := ReadCookie()

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: session_name, Value: session_value})

	fmt.Println("Fetching", url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
