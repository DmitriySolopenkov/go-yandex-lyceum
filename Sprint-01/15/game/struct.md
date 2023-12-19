- game
	- cmd // для хранения пакетов main
		-life
			- main.go // точка входа в программу
	- http
		- server // http сервер
			server.go // код сервера
			- handler  // регистрация функций обработчиков
				handler.go
	- internal
		- application  // конфигурация и код вызова приложения
			application.go
		- service // сервис, который инициализирует и хранит состояния игры
			service.go
	- pkg // для хранения пакетов
		- life // сама логика игры
			- life.go