package main

import (
	"encoding/json"
	"reflect"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
	var obj1 map[string]interface{}
	var obj2 map[string]interface{}

	err := json.Unmarshal(json1, &obj1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(json2, &obj2)
	if err != nil {
		return false, err
	}

	// Проверяем эквивалентность без учета порядка полей
	return reflect.DeepEqual(obj1, obj2), nil
}

// func main() {
// 	// Использование функции для сравнения
// 	jsonStr1 := []byte(`{"name":"John", "age":30, "car":null}`)
// 	jsonStr2 := []byte(`{"age":30, "name":"John", "car":null}`)

// 	equal, err := CompareJSON(jsonStr1, jsonStr2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if equal {
// 		println("JSON objects are equal")
// 	} else {
// 		println("JSON objects are not equal")
// 	}
// }

/*
JSON
Напишите функцию CompareJSON(json1, json2 []byte) (bool, error),
	которая принимает на входе два объекта json и сравнивает их.

	Если они равны, то функция должна вернуть true, nil, иначе false, nil, либо описание ошибки.

Примечания
порядок следования полей в json в равных объектах может быть разный
json не содержит вложенных объектов
*/
