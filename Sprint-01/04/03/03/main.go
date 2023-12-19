package main

import (
	"fmt"
	"net/http"
)

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/answer/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The answer is 42")
}

func requireBasicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			w.WriteHeader(403)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "userid" || pass != "password" {
			w.WriteHeader(401)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/answer/", handleAnswer)
	http.ListenAndServe(":5000", requireBasicAuth(http.DefaultServeMux))
}
