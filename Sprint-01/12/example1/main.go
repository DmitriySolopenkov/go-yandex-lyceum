package main

import (
	"fmt"
	"time"
)

func main() {
	// Записываем текущее время до начала операции
	startTime := time.Now()

	// Здесь вы можете выполнять какую-нибудь долгую операцию —
	// например, обработку большого объема данных или вычисления

	// Допустим, мы выполняем операцию, которая имитирует задержку
	for i := 0; i < 1000000; i++ {
		// Делаем что-то
	}

	// Записываем время после того, как операция выполнена
	elapsed := time.Since(startTime)

	// Проверяем, сколько времени заняла операция
	// Если операция заняла более 1 секунды, мы считаем её слишком долгой
	if elapsed.Seconds() > 1 {
		fmt.Println("Операция заняла слишком много времени")
	} else {
		fmt.Println("Операция завершена за", elapsed)
	}
}
