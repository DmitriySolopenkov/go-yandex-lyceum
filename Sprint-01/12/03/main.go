package main

import "time"

func ParseStringToTime(dateString, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

/*
Считать время
Напишите функцию ParseStringToTime(dateString, format string) (time.Time, error),
	которая анализирует строку в определенном формате и преобразует ее в значение time.Time.

Функция должна принимать строку и строку формата,
	возвращая проанализированное значение time.Time.

Примечания
Пример работы функции:

dateString := "2023-10-23 02:41:49"
format := "2006-01-02 15:04:05"
result, err := ParseStringToTime(dateString, format)
if err != nil {
    panic(err)
}
fmt.Println(result) // Output: 2023-10-23 02:41:49 +0000 UTC
*/
