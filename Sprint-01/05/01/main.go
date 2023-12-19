package main

import (
	"io"
	"net/http"
)

var Err string = "Something went wrong..."

func SendHTTPRequest(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
HTTP Клиент
Вам необходимо написать функцию SendHTTPRequest(url string) (string, error),
	которая делает GET-запрос к заданному URL и возвращает тело ответа в виде строки.

Примечания
	Нужный url подставляется внутри теста.
	В случае, если произошла ошибка, необходимо возвращать ошибку: "Something went wrong..."

Пример ответа функции:

"Hello, World!\n"
*/
