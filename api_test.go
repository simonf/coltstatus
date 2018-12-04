package coltstatus

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func sendTwoHundred(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("")
	w.Write(buffer.Bytes())
}
func TestCheckAPI(t *testing.T) {
	path := "/"
	port := ":2111"

	srv := &http.Server{Addr: port}

	ph := http.HandlerFunc(sendTwoHundred)
	http.Handle(path, ph)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	defer srv.Shutdown(nil)

	testline := "http://127.0.0.1" + port + "/ 200"
	url, code, err := parseLine(testline)
	println(url)
	println(strconv.Itoa(code))
	if err != nil || code != 200 {
		t.Fail()
	}
}
