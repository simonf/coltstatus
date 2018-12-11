//Package coltstatus checks whether a list of HTTP GET endpoints can be reached.
// It returns 200 if everything is OK
package coltstatus

import (
	"log"
	"net/http"
	"time"
)

// ApiTarget specifies the URL to call using a GET request and the HTTP status code to expect.
type ApiTarget struct {
	url             string
	expected_status int
}

// CheckDependentServices reads the specified file for a list of ApiTargets.
// Then it calls those targets and checks for the HTTP status code response.
// It returns:
// - 200 if all targets respond as expected
// - 500 as soon as a target fails to respond or responds with an unexpected HTTP status code
// - 404 if the config file cannot be found
func CheckDependentServices(targets []ApiTarget) int {
	c := make(chan bool)

	for _, element := range targets {
		go func() { c <- isAPIOK(element) }()
	}

	for _, _ = range targets {
		if !<-c {
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
