package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode"
)

func IsNoLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return tr
}

type Greeting struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type RPCResponse struct {
	Status string   `json:"greetings"`
	Result Greeting `json:"result"`
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", RPC(Sanitize(SetDefaultName(HelloHandler))))
	http.ListenAndServe(":8080", mux)
}

func RPC(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				resp := RPCResponse{
					Status: "error",
					Result: Greeting{},
				}
				jsonResp, err := json.Marshal(resp)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")

				w.Write(jsonResp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// обработку некорректного значения name
func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			// fmt.Fprint(w, "Hello, stranger!")
			name = "stranger"
			// return
		}
		w.Header().Set("Content-Type", "application/json")

		resp := map[string]interface{}{
			"greetings": "hello",
			"name":      name,
			// Здесь можно добавить другие поля
		}
		b, err := json.Marshal(resp)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, string(b))
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

/*
Возьмите сервер, отвечающий hello stranger из этого урока. Переделайте формат ответа на JSON вида:

{"greetings": "hello", "name": "stranger"}

При этом name берётся из запроса, а логика его подстановки не меняется.

Добавьте к этому middleware RPC(http.HandlerFunc), которая заменяет ответ на формат:

{"status": "ok", "result": {"greetings": "hello", "name": "stranger"}}

Так же, переделайте Middleware Sanitize, чтобы она возвращала panic в случае некорректного имени,
	и добавьте обработку этой паники в новой middleware RPC так, чтобы отдавать пользователю ответ:

{"status": "error", "result": {}}

Примечания
Вот примеры запросов и соответствующих ответов:

Запрос: GET /hello?name=Alice Ответ:
{
    "greetings": "Hello",
    "name": "Alice"
}

Запрос: GET /hello?name="" Ответ: Ошибка 500, так как имя не соответствует требованиям middleware Sanitize.

Ответ:

{
    "status": "error",
    "result": {}
}
*/
