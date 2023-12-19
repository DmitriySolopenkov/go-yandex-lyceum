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

func nameHandler(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("name")
	switch {
	case n == "":
		fmt.Fprint(w, "hello stranger")
	case len(n) > 0 && IsNoLetter(n):
		fmt.Fprint(w, "hello dirty hacker")
	case len(n) > 0:
		fmt.Fprint(w, "hello ", n)
	}
}

func main() {
	http.HandleFunc("/", nameHandler)
	http.ListenAndServe(":8080", nil)
}

/*
Напишите веб-сервер,
	который при обращении к нему будет возвращать приветствие с именем пользователя,
	полученным из параметра запроса.

Если параметр пустой или отсутствует,
	сервер должен вернуть приветствие hello stranger.

Если ответ содержит не только английские буквы,
	приветствие должно быть hello dirty hacker.

Веб-сервер должен быть запущен на порту с номером 8080.

Формат ввода
Пример запроса:

curl localhost:8080/?name=John
# hello John


curl localhost:8080
# hello stranger
Примечания
Для проверки имени проще всего применить одну из функций пакета strings.


*/
