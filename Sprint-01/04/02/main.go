package main

import (
	"fmt"
	"net/http"
)

var requestCount int

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// func MetricsHandler(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "rpc_duration_milliseconds_count ", requestCount)
// 		next.ServeHTTP(w, r)
// 	}
// }

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "rpc_duration_milliseconds_count ", requestCount)
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Fib(requestCount))
	requestCount++
}

// func MetricsHandler(w http.ResponseWriter, r *http.Request) {}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, Fib(Count))
// 	Count++
// }

// func main() {
// 	http.HandleFunc("/", helloHandler)
// 	http.ListenAndServe(":8080", nil)
// }

/*
Продолжаем развивать наши skill'ы!

Возьмите сервер Фибоначчи с метриками из предыдущего урока.

Переделайте увеличение метрики числа запросов в middleware Metrics(http.HandlerFunc)
	и добавьте этот middleware для url /.

Примечания
Функцию main создавать не надо.
Нумерация элементов в программировании всегда начинается с 0.

Для запросов к / (Фибоначчи):
Ответ с числом Фибоначчи, например:
Запрос: GET /
Ответ: "5".
Для запросов к /metrics (Метрики):
Ответ с метриками сервера, включая количество запросов (число запросов, увеличивающееся с каждым запросом), например:
Запрос: GET /metrics
Ответ:
rpc_duration_milliseconds_count 5
Где "5" - это количество запросов, которые были обработаны сервером.
Важно:
Ваша программа должна содержать следующий код (используйте как шаблон для своей программы):

package main
import (
"net/http"
)
var requestCount int
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
requestCount++
}
func MetricsHandler(w http.ResponseWriter, r *http.Request) { }

*/
