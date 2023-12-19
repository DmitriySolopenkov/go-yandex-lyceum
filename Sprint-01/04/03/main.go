package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))
	http.ListenAndServe(":5000", nil)
}

// Middleware функция для проверки заголовка Authorization
// func Authorization(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Получение значения заголовка Authorization
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			// Если заголовок не предоставлен, возвращаем статус 401 Unauthorized
// 			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized"))
// 			return
// 		}

// 		// Проверка корректности заголовка Authorization
// 		username, password, ok := parseBasicAuth(authHeader)
// 		if !ok || username != "userid" || password != "password" {
// 			// Если заголовок содержит некорректные данные, возвращаем статус 401 Unauthorized
// 			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized"))
// 			return
// 		}

// 		// Если авторизация успешна, передаем управление следующему обработчику
// 		next.ServeHTTP(w, r)
// 	}
// }

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			// expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			// expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

// Обработчик для пути /answer/
func answerHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем ответ "The answer is 42"
	w.Write([]byte("The answer is 42"))
}

// Функция для разбора значения заголовка Authorization в формате Basic
// func parseBasicAuth(authHeader string) (username, password string, ok bool) {
// 	// Проверяем префикс "Basic "
// 	if len(authHeader) > 6 && authHeader[:6] == "Basic " {
// 		// Декодируем Base64 и получаем значение формата "username:password"
// 		decoded, err := base64.StdEncoding.DecodeString(authHeader[6:])
// 		if err == nil {
// 			// Разделяем username и password
// 			credentials := string(decoded)
// 			colonIndex := -1
// 			for i, char := range credentials {
// 				if char == ':' {
// 					colonIndex = i
// 					break
// 				}
// 			}
// 			if colonIndex != -1 {
// 				return credentials[:colonIndex], credentials[colonIndex+1:], true
// 			}
// 		}
// 	}

//		return "", "", false
//	}
// func basicAuth(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Extract the username and password from the request
//         // Authorization header. If no Authentication header is present
//         // or the header value is invalid, then the 'ok' return value
//         // will be false.
// 		username, password, ok := r.BasicAuth()
// 		if ok {
//             // Calculate SHA-256 hashes for the provided and expected
//             // usernames and passwords.
// 			usernameHash := sha256.Sum256([]byte(username))
// 			passwordHash := sha256.Sum256([]byte(password))
// 			expectedUsernameHash := sha256.Sum256([]byte("your expected username"))
// 			expectedPasswordHash := sha256.Sum256([]byte("your expected password"))

//             // Use the subtle.ConstantTimeCompare() function to check if
//             // the provided username and password hashes equal the
//             // expected username and password hashes. ConstantTimeCompare
//             // will return 1 if the values are equal, or 0 otherwise.
//             // Importantly, we should to do the work to evaluate both the
//             // username and password before checking the return values to
//             // avoid leaking information.
// 			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
// 			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

//             // If the username and password are correct, then call
//             // the next handler in the chain. Make sure to return
//             // afterwards, so that none of the code below is run.
// 			if usernameMatch && passwordMatch {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}

//         // If the Authentication header is not present, is invalid, or the
//         // username or password is wrong, then set a WWW-Authenticate
//         // header to inform the client that we expect them to use basic
//         // authentication and send a 401 Unauthorized response.
// 		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 	})
// }

/*
Базовая HTTP-авторизация - это способ защитить веб-страницы или другие ресурсы от несанкционированного доступа.

Когда вы пытаетесь получить доступ к защищенному ресурсу, сервер запрашивает имя пользователя и пароль. Если они корректные, то вы получаете доступ к ресурсу.

Например, если у вас есть сайт с персональными данными, то базовая HTTP-авторизация поможет защитить эти данные от злоумышленников.

При использовании базовой авторизации поверх HTTP используется заголовок Authorization

Напишите веб-сервер с использованием базовой HTTP-авторизации на пути /answer/.

Сервер должен проверять наличие и корректность заголовка Authorization и возвращать ответ The answer is 42 при успешной авторизации.

При запросе без заголовка Authorization сервер должен вернуть статус 401 Unauthorized и запросить авторизацию.

Используйте необходимые пакеты и функции для реализации данного функционала.

Middleware функцию назовите Authorization(http.HandlerFunc)

Примечания
Корректные данные для пользователя

Login: userid
Password: password
Пример ответа:

POST /api/users HTTP/1.1 # метод и URL
Host: example.com # обязательный заголовок Host
Content-Type: application/json  # заголовок с типом данных
Authorization: Basic userid:password

{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "123456"
} # тело запроса
Успешная авторизация:

curl -X GET http://127.0.0.1:5000/answer/ -H "Authorization: Basic YWRtaW46cGFzc3dvcmQ="
Ожидаемый ответ:
The answer is 42
Отсутствие или некорректная авторизация: Пример запроса (с использованием cURL без заголовка Authorization):

curl -X GET http://127.0.0.1:5000/answer/
Ожидаемый ответ:
Unauthorized

Пример запроса (с использованием cURL с некорректными данными):
curl -X GET http://127.0.0.1:5000/answer/ -H "Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ="
Ожидаемый ответ:
Unauthorized
*/
