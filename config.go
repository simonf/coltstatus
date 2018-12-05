package coltstatus

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadConfigFile reads the specified file and returns a list of ApiTargets.
// If the file does not exist it returns an error
// The file should be structured with one line per target containing:
// - the URL to call using an HTTP GET request
// - a space
// - the expected HTTP status code
//
// e.g.: http://myserver:1234/api/hello 200
//
// Lines starting with a # symbol will be ignored
func ReadConfigFile(cfg_file string) ([]ApiTarget, error) {
	rv := make([]ApiTarget, 0)
	f, err := os.Open(cfg_file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		if !shouldIgnore(line) {
			url, status, err := parseLine(line)
			if err != nil {
				return nil, err
			}
			tocheck := ApiTarget{url: url, expected_status: status}
			rv = append(rv, tocheck)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading ", cfg_file, ": ", err)
		return nil, err
	}
	return rv, nil
}

func shouldIgnore(line string) bool {
	if strings.HasPrefix(line, "#") || len(line) < 11 {
		return true
	}
	return false
}

// Return a tuple of URL + expected HTTP status code
func parseLine(line string) (string, int, error) {
	words := strings.Fields(line)
	if len(words) < 2 {
		ertxt := fmt.Sprintf("Missing content check in line: %s", line)
		log.Print(ertxt)
		return "", 0, errors.New("Missing content")
	}
	url := words[0]
	code, err := strconv.ParseInt(strings.TrimSpace(line[len(url):]), 10, 0)
	if err != nil {
		log.Print(fmt.Sprintf("Could not find status code as second field in %s", line))
		return "", 0, err
	} else {
		return url, int(code), nil
	}
}
