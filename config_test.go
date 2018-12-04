package coltstatus

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestParseLine(t *testing.T) {
	testline := "http://172.22.0.3:5000/ 200"
	url, code, err := parseLine(testline)
	println(url)
	println(strconv.Itoa(code))
	if err != nil {
		t.Fail()
	}
}

func TestIgnoreComment(t *testing.T) {
	testline := "#this should be ignored"
	b := shouldIgnore(testline)
	if !b {
		t.Fail()
	}
	okline := "Should be ok"
	b = shouldIgnore(okline)
	if b {
		t.Fail()
	}
}
func TestReadConfigFile(t *testing.T) {
	fname := "test.txt"
	var buffer bytes.Buffer
	buffer.WriteString("http://172.22.0.3:5000/ 200")
	err := ioutil.WriteFile(fname, buffer.Bytes(), 0644)
	defer os.Remove(fname)
	if err != nil {
		t.Fail()
	}
	// fname := "apis.txt"
	targets := readConfigFile(fname)
	if len(targets) < 1 {
		t.Fail()
	}

}
