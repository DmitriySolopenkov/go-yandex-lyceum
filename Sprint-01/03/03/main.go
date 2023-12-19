package main

import (
	"fmt"
	"net/http"
)

var Count int

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

type countHandler struct{}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "rpc_duration_milliseconds_count ", Count)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/metrics/", &countHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprint(w, Fib(Count))
		Count++
	})
	http.ListenAndServe(":8080", mux)
}

/*
Напишите веб-сервер, который будет считать метрики времени ответа сервиса.

Возьмите в качестве основы веб-сервер из предыдущей задачи, вычисляющий числа Фибоначчи, и добавьте к нему хендлер /metrics, который отдаёт:

rpc_duration_milliseconds_count 10
где 10 — число раз, которое вызвали хендлер, отвечающий числами Фибоначчи. Выглядит это так:

curl http://localhost:8080/metrics

# rpc_duration_milliseconds_count 0

curl http://localhost:8080/

# 0

curl http://localhost:8080/metrics

# rpc_duration_milliseconds_count 1

# ...

curl http://localhost:8080/

# 34

curl http://localhost:8080/metrics

# rpc_duration_milliseconds_count 10
*/
