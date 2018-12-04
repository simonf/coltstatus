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

func readConfigFile(cfg_file string) []ApiTarget {
	rv := make([]ApiTarget, 0)
	f, err := os.Open(cfg_file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		if !shouldIgnore(line) {
			url, status, err := parseLine(line)
			if err != nil {
				log.Fatal(err)
			}
			tocheck := ApiTarget{url: url, expected_status: status}
			rv = append(rv, tocheck)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading ", cfg_file, ": ", err)
		os.Exit(1)
	}
	return rv
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
