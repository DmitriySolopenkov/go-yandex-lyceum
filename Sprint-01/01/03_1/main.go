package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(data []byte) (map[string][]byte, error) {
	students := make(map[string][]Student)

	var studentsArray []Student
	err := json.Unmarshal(data, &studentsArray)
	if err != nil {
		return nil, err
	}

	for _, student := range studentsArray {
		class := student.Class
		name := student.Name

		students[class] = append(students[class], Student{Class: class, Name: name})
	}

	jsonData := make(map[string][]byte)
	for class, studentsGroup := range students {
		groupJSON, err := json.Marshal(studentsGroup)
		if err != nil {
			return nil, err
		}

		jsonData[class] = groupJSON
	}
	return jsonData, nil
}

func main() {
	input := []byte(`[
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
	data, err := splitJSONByClass(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
