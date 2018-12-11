package main

import (
	"bytes"
	"fmt"
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
	cfgfile := "apis.txt"

	if len(args) > 0 {
		port, _ = strconv.Atoi(args[0])
	}

	if len(args) > 1 {
		cfgfile = args[1]
	}

	_, err := coltstatus.ReadConfigFile(cfgfile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Config file not found. Exiting.", cfgfile, ": ", err)
		os.Exit(1)
	}

	startServer(path, port)
}

func startServer(path string, port int) {
	log.Print(" Listening on port " + strconv.Itoa(port) + " for path " + path)
	ph := http.HandlerFunc(getRequestHandler)
	http.Handle(path, ph)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	// logRequest(r)
	targets, err := coltstatus.ReadConfigFile("apis.txt")
	if err != nil {
		w.WriteHeader(404)
		return
	}
	result := coltstatus.CheckDependentServices(targets)
	var buffer bytes.Buffer
	w.WriteHeader(result)
	buffer.WriteString("")
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
