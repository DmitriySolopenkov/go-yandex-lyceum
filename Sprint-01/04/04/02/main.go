package main

import (
	"encoding/json"
	"net/http"
	"unicode"
)

type Greeting struct {
	Greetings string `json:"greetings,omitempty"`
	Name      string `json:"name,omitempty"`
}

type RPCResponse struct {
	Status string   `json:"status"`
	Result Greeting `json:"result,omitempty"`
}

// type RPCResponseInvalid struct {
// 	Status string                 `json:"status"`
// 	Result map[string]interface{} `json:"result"`
// }

func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	res := RPCResponse{
		Status: "ok",
		Result: Greeting{
			Greetings: "hello",
			Name:      name,
		},
	}

	resp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
			res := RPCResponse{
				Status: "ok",
				Result: Greeting{
					Greetings: "hello",
					Name:      name,
				},
			}

			resp, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.Write(resp)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// -- FAIL: TestHelloHandler (0.00s)
// --- FAIL: TestHelloHandler/Invalid_Name (0.00s)
// source_test.go:39: Expected response body "{"status":"error","result":{}}", but got "{"result":{},"status":"error"}"
// FAIL

func RPC(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				res := RPCResponse{
					Status: "error",
				}
				resp, err := json.Marshal(res)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")

				w.Write(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !isValidName(name) {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)

	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", RPC(Sanitize(SetDefaultName(HelloHandler))))
	http.ListenAndServe(":8080", mux)
}

// ################

/*
-- FAIL: TestHelloHandler (0.00s)
    --- FAIL: TestHelloHandler/Invalid_Name (0.00s)
        source_test.go:39: Expected response body "{"status":"error","result":{}}", but got "{"status":"error","result":{"greetings":"","name":""}}"
FAIL
*/
