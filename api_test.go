package coltstatus

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

type testServer struct {
	port string
	code int
}

// Start an HTTP server and return it so it can be shut down
func oneShotServer(port string, code int) *http.Server {
	//	path := "/"

	ph := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buffer bytes.Buffer
		buffer.WriteString("")
		w.WriteHeader(code)
		w.Write(buffer.Bytes())
	})

	srv := &http.Server{Addr: ":" + port, Handler: ph}

	// http.Handle(path, ph)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	return srv
}

// Make an array of HTTP servers, on specified ports and returning specified HTTP status codes
func makeServers(pl []testServer) []*http.Server {
	retval := make([]*http.Server, len(pl))
	for i, ts := range pl {
		retval[i] = oneShotServer(ts.port, ts.code)
	}
	return retval
}

func TestCheckAPI(t *testing.T) {
	tests := make([]ApiTarget, 3)
	servers := make([]testServer, 3)

	tests[0] = ApiTarget{url: "http://localhost:2111", expected_status: 200}
	tests[1] = ApiTarget{url: "http://localhost:2112", expected_status: 200}
	tests[2] = ApiTarget{url: "http://localhost:2113", expected_status: 500}
	servers[0] = testServer{port: "2111", code: 200}
	servers[1] = testServer{port: "2112", code: 200}
	servers[2] = testServer{port: "2113", code: 500}

	for _, srv := range makeServers(servers) {
		defer srv.Shutdown(nil)
	}

	if 200 != CheckDependentServices(tests) {
		t.Fail()
	}

	tests[2] = ApiTarget{url: "http://localhost:2113", expected_status: 200}
	if 500 != CheckDependentServices(tests) {
		t.Fail()
	}

}
