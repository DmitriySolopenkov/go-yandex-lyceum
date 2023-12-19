package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

// Middleware для проверки и нормализации имени
func Sanitize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Invalid name", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Обработчик для запросов "/hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	greetings := "Hello"
	response := map[string]string{"greetings": greetings, "name": name}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Middleware для обработки ответа
func RPC(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				response := map[string]interface{}{"status": "error", "result": map[string]interface{}{}}
				jsonResponse, err := json.Marshal(response)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResponse)
			}
		}()
		next.ServeHTTP(w, r)
		response := map[string]interface{}{"status": "ok", "result": map[string]interface{}{}}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})
}

func main() {
	// r := http.NewServeMux()

	// Используем middleware Sanitize для обработки запросов
	// r.HandleFunc("/hello", HelloHandler).Methods("GET").Name("Hello").Handler(Sanitize)
	mux := http.NewServeMux()
	mux.Handle("/", Sanitize(RPC(HelloHandler())))

	// Используем middleware RPC для обработки ответов
	http.Handle("/", RPC(r))

	fmt.Println("Server is running on :8080")

	http.ListenAndServe(":8080", r)
}
