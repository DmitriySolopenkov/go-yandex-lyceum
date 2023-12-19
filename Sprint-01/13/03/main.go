package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Создаем маршрутизатор HTTP
	router := http.NewServeMux()

	// Обработчик для всех запросов
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Записываем информацию о запросе в лог
		log.Printf("%s %s %s %s %s\n", time.Now().Format("2006/01/02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr, r.Header)

		// Отправляем ответ клиенту
		fmt.Fprintf(w, "Hello, World!")
	})

	// Запускаем веб-сервер на порту 8080
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

/*
Логгирование HTTP сервера
Напишите веб сервер,
	который логгирует информацию о запросе вида
		2023/11/01 09:41:19 GET / [::1]:55250 {} 7.708µs 2023/11/01 09:41:19 GET / [::1]:55250 {} 1.166µs

		В функции main должен быть запуск веб-сервера на порту 8080
*/
