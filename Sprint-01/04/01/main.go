package main

import (
	"fmt"
	"net/http"
	"unicode"
)

func IsNoLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// обработку некорректного значения name
func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprint(w, "Hello, stranger!")
			name = "stranger"
			return
		}
		next.ServeHTTP(w, r)
	}
}

// обработку пустого значения
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsNoLetter(r.URL.Query().Get("name")) {
			fmt.Fprint(w, "Hello, dirty hacker!")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name, "!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", Sanitize(SetDefaultName(HelloHandler)))
	http.ListenAndServe(":8080", mux)
}

/*
Давайте модифицируем самый первый сервер, который мы сделали.

Да-да, тот самый, что отвечает hello stranger.

Переделайте обработку некорректного значения name в middleware Sanitize(http.HandlerFunc)
	и обработку пустого значения в middleware SetDefaultName(http.HandlerFunc).

Примечания
Обработка некорректного значения означает, что name не должно содержать символов,
	которые не подходят для имени.

То есть, оно не должно содержать символов !@#$%^&*()_+.

При получении в поле name пустой строки метод SetDefaultName
	должен устанавливать поле name равным "stranger".
*/
