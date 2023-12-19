package main

import (
	"fmt"
	"log"
	"net/http"
)

var users = map[string]string{
	"test": "secret",
}

func isAuthorised(name, password string) bool {
	pass, ok := users[name]
	if !ok {
		return false
	}

	return password == pass
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	name, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Unauthorized"}`))
		return
	}

	if !isAuthorised(name, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Unauthorized"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "The answer is 42"}`))
	// return
}

func main() {
	http.HandleFunc("/answer/", Authorization)
	fmt.Println("Starting Server at port :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
