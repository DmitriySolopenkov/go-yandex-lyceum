package main

import (
	"context"
	"io"
	"net/http"
)

var Err string = "Something went wrong..."

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Err, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Err, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Err, err
	}

	return string(body), nil
}

/*
HTTP Клиент с контекстом
Используя http.Client и context,
	напишите функцию SendHTTPRequestWithContext(ctx context.Context, url string) (string, error),
	которая делает GET-запрос к заданному URL и принимает контекст для управления запросом.

Данная задача похожа на предыдущую, однако следует добавить контекст в функцию,
	которая выполняет запрос. Ошибки обрабатывайте аналогично.

Для выполнения запроса используйте функцию NewRequestWithContext.

Примечания
Пример возвращаемой функцией строки:

"Hello, World!\n"
*/
