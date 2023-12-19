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
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "hello stranger" {
		t.Fatal("incorrect answer", string(body), "expect 'hello stranger'")
	}
}

func TestName(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/?name=Vasya")
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "hello Vasya" {
		t.Fatal("incorrect answer", string(body), "expect 'hello Vasya'")
	}
}

func TestHacker(t *testing.T) {
	resp, err := http.Get("http://localhost:
