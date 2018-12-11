package coltstatus

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseLine(t *testing.T) {
	testline := "http://172.22.0.3:5000/ 200"
	_, _, err := parseLine(testline)
	// println(url)
	// println(strconv.Itoa(code))
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

func TestReadEmptyConfigFile(t *testing.T) {
	fname := "empty.txt"
	var buffer bytes.Buffer
	buffer.WriteString("")
	err := ioutil.WriteFile(fname, buffer.Bytes(), 0644)
	defer os.Remove(fname)
	if err != nil {
		t.Fail()
	}
	// fname := "apis.txt"
	targets, err := ReadConfigFile(fname)
	if err != nil || len(targets) != 0 {
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
	targets, _ := ReadConfigFile(fname)
	if len(targets) < 1 {
		t.Fail()
	}

}
