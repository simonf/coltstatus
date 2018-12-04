package coltstatus

import (
	"log"
	"net/http"
	"time"
)

type ApiTarget struct {
	url             string
	expected_status int
}

func CheckDependentServices(filename string) int {
	targets := readConfigFile(filename)
	for _, element := range targets {
		if !isAPIOK(element) {
			return 500
		}
	}
	return 200
}

func isAPIOK(api ApiTarget) bool {
	log.Print(api.url)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(api.url)
	if err != nil || resp.StatusCode != api.expected_status {
		return false
	}
	return true
}
