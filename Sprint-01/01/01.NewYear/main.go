package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student

	// Распаковываем входные данные JSON в структуру
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		panic(err)
	}
	// fmt.Println(students)
	// Модифицируем каждого студента, добавляя 1 к полю Grade
	for i := range students {
		fmt.Println(students[i].Name)
		fmt.Println(students[i].Grade)
		students[i].Grade++
	}

	// Конвертируем обновленную структуру обратно в JSON
	updatedJSON, err := json.Marshal(students)
	if err != nil {
		return nil, err
	}

	return updatedJSON, nil
}

func main() {
	students := []byte(`[
		{
			"name": "Oleg",
			"grade": 9
		},
		{
			"name": "Katya",
			"grade": 10
		}
	]`)
	updatedJSON, err := modifyJSON(students)
	if err != nil {
		// Обработка ошибки
	} else {
		// Вывод обновленных данных
		fmt.Println(string(updatedJSON))
	}

}

/*
С новым учебным годом!
1 сентября каждого учебного года во всех базах данных школьников происходит великий пересчёт.
Напишите функцию modifyJSON(jsonData []byte) ([]byte, error),
	которая принимает данные в формате JSON,
	добавляет 1 год к полю grade и возвращает обновлённые данные также в формате JSON.

Формат ввода
[
    {
        "name": "Oleg",
        "grade": 9
    },
    {
        "name": "Katya",
        "grade": 10
    }
]
Формат вывода
[
    {
        "name": "Oleg",
        "grade": 10
    },
    {
        "name": "Katya",
        "grade": 11
    }
]
Примечания
Структура ученика

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}
*/
