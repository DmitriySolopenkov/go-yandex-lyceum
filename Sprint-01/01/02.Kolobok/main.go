package main

import (
	"encoding/json"
	"fmt"
)

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var mergedData []map[string]interface{}

	// Добавляем всё из вводных данных в map
	for _, jsonData := range jsonDataList {
		var data []map[string]interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			panic(err)
		}
		mergedData = append(mergedData, data...)
	}

	// Конвертируем обновленную структуру обратно в JSON
	mergedJSON, err := json.Marshal(mergedData)
	if err != nil {
		return nil, err
	}

	return mergedJSON, nil
}

func main() {
	student1 := []byte(`[
		{
			"name": "Oleg",
			"class": "9B"
		},
		{
			"name": "Ivan",
			"class": "9A"
		}
	]`)
	student2 := []byte(`[
		{
			"name": "Maria",
			"class": "9B"
		},
		{
			"name": "John",
			"class": "9A"
		}
	]`)
	mergedJSON, err := mergeJSONData(student1, student2)
	if err != nil {
		// Обработка ошибки
	} else {
		// Вывод обновленных данных
		fmt.Println(string(mergedJSON))
	}
}

/*
По амбарам помела, по сусекам поскребла — базу данных получила. Напишите функцию mergeJSONData(jsonDataList ...[]byte) ([]byte, error), которая принимает несколько экземпляров данных в формате JSON, объединяет их в один экземпляр и возвращает его.

Примечания
Например: В функцию передаются два JSON:

[
    {
        "name": "Oleg",
        "class": "9B"
    },
    {
        "name": "Ivan",
        "class": "9A"
    }
]
и

[
    {
        "name": "Maria",
        "class": "9B"
    },
    {
        "name": "John",
        "class": "9A"
    }
]
На выходе нужно получить:

[
    {
        "class": "9B",
        "name": "Oleg"
    },
    {
        "class": "9A",
        "name": "Ivan"
    },
    {
        "class": "9B",
        "name": "Maria"
    },
    {
        "class": "9A",
        "name": "John"
    }
]

*/
