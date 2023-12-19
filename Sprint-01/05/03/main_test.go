package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(time.Second)
	os.Exit(m.Run())
}

func TestDefault(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		return
	}
	request.Header.Set("X-Request-ID", "123456789")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "Hello! RequestID: 123456789" {
		t.Fatal("incorrect answer", string(body), "expect 'Hello! RequestID: 123456789'")
	}
}

func TestNoRequestID(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "Hello! RequestID not found" {
		t.Fatal("incorrect answer", string(body), "expect 'Hello! RequestID not found'")
	}
}

// --- FAIL: TestNoRequestID (0.00s)
//     source_test.go:57: incorrect answer Hello! RequestID: RequestID not found expect Hello! RequestID not found
// FAIL
