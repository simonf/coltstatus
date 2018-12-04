package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"

	"simonf.net/coltstatus"
)

func main() {
	args := os.Args[1:]
	port := 8111
	path := "/"

	if len(args) > 0 {
		port, _ = strconv.Atoi(args[0])
	}

	StartServer(path, port)
}

func StartServer(path string, port int) {
	log.Print(" Listening on port " + strconv.Itoa(port) + " for path " + path)
	ph := http.HandlerFunc(getRequestHandler)
	http.Handle(path, ph)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	result := coltstatus.CheckDependentServices("apis.txt")
	var buffer bytes.Buffer
	w.WriteHeader(result)
	buffer.WriteString("hi")
	w.Write(buffer.Bytes())
}

func logRequest(r *http.Request) {
	// log everything we can find in the request:
	log.Print("Method: " + r.Method)
	log.Print("URL: " + r.URL.String())
	log.Print("Proto: " + r.Proto)
	log.Print("Headers:")
	log.Print(r.Header)
	log.Print("ContentLength: " + strconv.Itoa(int(r.ContentLength)))
	log.Print("TransferEncoding")
	log.Print(r.TransferEncoding)
}
