package main

import (
	"context"
	"fmt"
	"net/http"
)

type key int

const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("X-Request-ID")
	return context.WithValue(ctx, requestIDKey, reqID)
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	reqID := requestIDFromContext(r.Context())
	if reqID == "" {
		fmt.Fprint(w, "Hello! RequestID not found")
	} else {
		fmt.Fprintf(w, "Hello! RequestID: %s", reqID)
	}
}

func main() {
	http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))
	http.ListenAndServe(":8080", nil)
}

/*
Middleware Context
Вам необходимо создать веб-сервер с
	Middleware RequestIDMiddleware(next http.Handler) http.Handler
	для HTTP-обработчика HelloHandler(w http.ResponseWriter, r *http.Request),
	который будет добавлять информацию из заголовка "X-Request-ID" в контекст запроса
	и затем использовать эту информацию в самом обработчике. Если "X-Request-ID" не передается -
	необходимо написать об этом в формате в виде "RequestID not found".

Не забудьте про функцию main, которая должна содержать методы для запуска сервера в виде

func main() {
    http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

    http.ListenAndServe(":8080", nil)
}
Примечания
Запрос:

HTTP метод: GET
Путь: "/"
Заголовок: X-Request-ID: 12345
Ожидаемый ответ:

Статус: 200 OK
Тело ответа: "Hello! RequestID: 12345"
Запрос:

HTTP метод: GET
Путь: "/"
Ожидаемый ответ:

Статус: 200 OK
Тело ответа: "Hello! RequestID not found"
*/
