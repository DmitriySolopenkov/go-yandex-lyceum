package main

import (
	"encoding/json"
	"errors"
)

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var data []map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	result := make(map[string][]byte)
	for _, val := range data {
		var class string
		class, ok := val["class"].(string)
		if !ok {
			return nil, errors.New("Invalid data format: 'class' field is missing or not a string")
		}

		// Convert the val to JSON
		itemJSON, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		if existingData, exists := result[class]; exists {
			// If the class already exists in the result, append the val JSON
			result[class] = append(existingData, itemJSON...)
		} else {
			// If the class doesn't exist in the result, create a new entry
			// result[class] = []byte("[" + string(itemJSON) + "]")
			result[class] = []byte(string(itemJSON))
		}
	}

	return result, nil
}

func main() {
	jsonData := []byte(`[
		{
			"name": "Oleg",
			"class": "9B"
		},
		{
			"name": "Ivan",
			"class": "9A"
		},
		{
			"name": "Maria",
			"class": "9B"
		},
		{
			"name": "John",
			"class": "9A"
		}
	]`)

	result, err := splitJSONByClass(jsonData)
	if err != nil {
		panic(err)
	}

	// for idx, val := range result {
	// 	println(idx, string(val))
	// }
	for _, val := range result {
		println(string(val))
	}
	// fmt.Println(result)

	// expectedJSONMap := map[string][]byte{
	// 	"9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
	// 	"9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]`),
	// }
	// fmt.Println(expectedJSONMap)
}

/*
На линейке ученикам нужно сгруппироваться по классам.

Проведём линейку для базы данных.

Напишите функцию splitJSONByClass(jsonData []byte) (map[string][]byte, error),
	которая принимает данные в формате JSON и возвращает мапу, в которой ключи — классы,
	а значения — данные в формате JSON, которые к этому классу относятся.

Примечания
Например: Входные данные

[
    {
        "name": "Oleg",
        "class": "9B"
    },
    {
        "name": "Ivan",
        "class": "9A"
    },
    {
        "name": "Maria",
        "class": "9B"
    },
    {
        "name": "John",
        "class": "9A"
    }
]
Выходные данные должны быть в виде map:

map[string][]byte{
        "9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
        "9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]`),
}

--- FAIL: TestSplitJSONByClass (0.00s)
    source_test.go:40: Expected JSON data for class 9A to be [{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}], but got [{"class":"9A","name":"Ivan"}]{"class":"9A","name":"John"}
    source_test.go:40: Expected JSON data for class 9B to be [{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}], but got [{"class":"9B","name":"Oleg"}]{"class":"9B","name":"Maria"}
FAIL
*/
