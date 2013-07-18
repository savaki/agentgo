package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetAmazonDataTimeOut(t *testing.T) {
	start := time.Now()
	data, err := getAmazonUserData()
	if data != nil {
		t.Fatalf("expected data to be nil since we're testing outside amazon; actual was %+v\n", data)
	}
	elapsed := time.Since(start)
	if elapsed < TIMEOUT {
		t.Fatalf("expected timeout to be at least %v; actual was %v\n", TIMEOUT, elapsed)
	}
	fmt.Println(err)
}

type ClosingBuffer struct {
	*bytes.Buffer
}

func (*ClosingBuffer) Close() error {
	return nil
}

func NewClosingBuffer(obj interface{}) (io.ReadCloser, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return &ClosingBuffer{bytes.NewBuffer(data)}, nil
}

func TestGetUserData(t *testing.T) {
	// Given
	expected := &UserData{Server: "localhost", Port: "8080"}
	f := func() (io.ReadCloser, error) { return NewClosingBuffer(expected) } // wrap our RaedCloser in a func

	// When
	actual, err := getUserData(f)

	// Then
	if err != nil {
		t.Fatalf("expected to pass; actual was err %s\n", err)
	}
	if actual.Server != expected.Server {
		t.Fatalf("expected actual.Server to be %s; actual was %s\n", expected.Server, actual.Server)
	}
	if actual.Port != expected.Port {
		t.Fatalf("expected actual.Port to be %s; actual was %s\n", expected.Port, actual.Port)
	}
}

func TestWriteUserAgentConfig(t *testing.T) {
	// Given
	sample := `
GO_SERVER=127.0.0.1
export GO_SERVER
GO_SERVER_PORT=8153
export GO_SERVER_PORT
AGENT_WORK_DIR=/var/lib/go-agent
export AGENT_WORK_DIR
DAEMON=Y
VNC=N
JAVA_HOME=/usr/lib/jvm/java-7-oracle/jre
export JAVA_HOME

export GOPATH=${HOME}/gocode
export PATH=${PATH}:${GOPATH}/bin
`

	filename := "sample.txt"
	config := &UserData{Server: "localhost", Port: "8080"}
	ioutil.WriteFile(filename, []byte(sample), 0644)

	// When
	writeGoAgentConfig(config, filename)

	// Then
	var expected string
	actual, _ := ioutil.ReadFile(filename)

	expected = GO_SERVER + config.Server
	if !strings.Contains(string(actual), expected) {
		t.Fatalf("expected config to contain string, %s\n", expected)
	}

	expected = GO_SERVER_PORT + config.Port
	if !strings.Contains(string(actual), expected) {
		t.Fatalf("expected config to contain string, %s\n", expected)
	}
}

func TestGoServer(t *testing.T) {
	text := `

GO_SERVER=1234

`
	parts := GO_SERVER_RE.FindAllString(text, -1)
	if len(parts) != 1 {
		t.Fatal("expected to find one match\n")
	}
	if parts[0] != "GO_SERVER=1234" {
		t.Fatalf("expected parts[0] to be GO_SERVER=1234; actual was %s\n", parts[0])
	}
}

func TestWriteToFile(t *testing.T) {
	// Given
	filename := "sample.err"
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("deleting file, %s\n", filename)
		os.Remove(filename)
	}
	expected := fmt.Sprintf("ugh!  something bad happened %d", time.Now().Unix())

	// When
	err := writeErrorMessage(errors.New(expected), filename)

	// Then
	if err != nil {
		panic(err)
	}
	actual, err := ioutil.ReadFile(filename) // ensure the file exists
	if err != nil {
		panic(err)
	}
	if string(actual) != expected {
		t.Fatalf("expected %s; actual was %s\n", expected, string(actual))
	}
}
