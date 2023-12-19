package main

import (
	"log"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	// Открываем файл
	file, err := os.OpenFile("temp_output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	// Закрываем файл после выхода из main
	defer file.Close()
	// Конфигурируем логгер, чтобы он выводил лог в файл
	log.SetOutput(file)
	log.Println("hello world")

	return nil
}

/*
Пишем в файл
Напишите функцию WriteToLogFile(message string, fileName string) error,
	которая пишет в файл output.txt строку "hello world"
*/
