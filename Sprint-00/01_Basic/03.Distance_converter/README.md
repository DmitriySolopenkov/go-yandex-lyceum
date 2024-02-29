# Конвертер расстояния

| Ограничения       |                   |
| -                 | -                 |
|Ограничение времени|1 секунда          |
|Ограничение памяти |64.0 Мб            |
|Ввод               |стандартный ввод   |
|Вывод              |стандартный вывод  |

У Ани отец – моряк, поэтому Аня хочет написать программу, которая переводит введённое пользователем число метров в морские мили и выводит результат на экран. Где-то возникла ошибка – помогите Ане.

```golang
package main

import "fmt"

func main() {
    meters = 0.0
    fmt.Scan(&meters)

    const exchangeRate = 1852
    seaMiles := meters / exchangeRate
    fmt.Println(seaMiles)
}
```

## Пример

|Ввод (тип)     |Вывод (тип)                    |
|-              |-                              |
|10 (float64)   |0.005399568034557235 (float64) |