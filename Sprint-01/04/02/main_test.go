package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciHandler(t *testing.T) {
	requestCount = 0

	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{"First Fibonacci Number", http.StatusOK, "0"},
		{"Second Fibonacci Number", http.StatusOK, "1"},
		{"Third Fibonacci Number", http.StatusOK, "1"},
		{"Fourth Fibonacci Number", http.StatusOK, "2"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(FibonacciHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rr.Code)
			}

			if rr.Body.String() != tc.expectedBody {
				t.Errorf("Expected body %s, but got %s", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
