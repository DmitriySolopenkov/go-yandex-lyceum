# Стандартные тесты в Go

## Введение в тестирование

> **Аннотация**
> Тесты — это программы, которые проверяют работу других программ. Тесты позволяют убедиться, что программа работает корректно и не содержит ошибок

Интересные факты:

- часто тесты требуют большего объёма работы нежели написание самой программы;
- в профессиональной разработке тесты иногда пишутся перед тем, как писать саму программу, это называется TDD (Test Driven Development);
- в профессиональной разработке тесты — обязательная часть работы программиста.

Для нашего удобства разработчики Go предоставляют стандартный тестовый фреймворк, который позволяет писать тесты.

В основном тесты пишутся, чтобы протестировать часть кода программы, например, функцию или метод. Такие тесты называют модульными тестами.

Стандартный тестовый фреймворк Go предоставляет многие удобства для написания тестов. Он позволяет запускать тесты параллельно, группировать тесты в подгруппы, пропускать тесты, если они не могут быть запущены, и многое другое.

## Соглашение о тестах

Стандартные тесты в Go пишутся с применением пакета testing. Тестовые функции должны начинаться с Test, сигнатура должна принимать один аргумент типа \*testing.T.

```go
func TestSomething(t *testing.T) {
// ...
}
```

Файл с тестами лучше называть так же, как и тестируемый файл, но с суффиксом **\_test**. Например, если мы тестируем файл **main.go**, то тесты должны быть в файле **main_test.go**. От правила с окончанием **\_test** отступить нельзя.

### Использование t.Run для запуска подтестов

Если ваш тест имеет несколько логических частей, вы можете использовать t.Run для создания подтестов. Это упрощает отслеживание ошибок и облегчает понимание, какая часть теста не прошла.

Пример:

```go
func TestConcurrentFunctionality(t *testing.T) {
t.Run("Increment", testIncrement)
t.Run("Decrement", testDecrement)
}

func testIncrement(t *testing.T) {
// тело теста инкрементации...
}

func testDecrement(t *testing.T) {
// тело теста декрементации...
}
```

## Запуск тестов. Команда go test

Для запуска тестов в пакете используется команда go test. Она принимает несколько флагов, которые позволяют конфигурировать запуск тестов.

Давайте попробуем выполнить эту команду в вашем проекте.

Создадим пакет printer с функцией, которая генерирует приветственное сообщение на основе имени человека. Создайте папку для проекта и добавьте файл main.go:

```go
package printer

import "fmt"

func PrintHello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
```

Затем мы пишем простой модульный тест для него. Добавим тесты в файл main_test.go:

```go
package printer

import "testing"

func TestPrintHello(t *testing.T) {
    got := PrintHello("Igor")
    expected := "Hello, Igor!"

    if got != expected {
        t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
    }
}
```

Инициализируем модуль

```bash
go mod init
```

Запустим тесты

```bash
go test
```

Скорее всего вы увидите что-то типа

```bash
go test
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.003s
```

## Флаг -v

Команда **go test** запускает тесты в текущем пакете. В нашем случае тесты прошли успешно. При этом мы не видим никакой дополнительной информации о тестах. Давайте попробуем запустить тесты с флагом **-v**:

```bash
go test -v
```

Флаг -v позволяет выводить подробную информацию о тестах. Теперь мы видим более подробную информацию о тестах:

```bash
   go test -v
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.001s
```

## Ошибка в работе теста

В отличие от обычных программ, модульные тесты при ошибке дают нам полезную информацию о том, что пошло не так. Не стоит расстриваться если тесты падают, это хорошо. Это значит, что тесты работают и выявляют ошибки в программе.

Давайте попробуем изменить тест так, чтобы он падал:

```go
func TestPrintHello(t *testing.T) {
    t.Fail()
}
```

Теперь, если мы запустим тесты

```bash
   go test -v
=== RUN   TestPrintHello
--- FAIL: TestPrintHello (0.00s)
FAIL
exit status 1
FAIL    github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.003s
```

Мы видим, что тест TestPrintHello выполнился с ошибкой.

Откатите изменения в тесте

## Кеширование результатов тестов

Кеширование результатов тестов было добавлено в Go 1.10. Идея проста — если мы запускаем тесты для пакета и они проходят, то при последующих запусках тех же тестов в "том же окружении" (определение следует далее) будут использоваться кешированные результаты вместо повторного запуска тестов.

Первый раз, когда мы запускаем тест, Go действительно выполняет его, что подтверждается продолжительностью времени, отображаемой рядом с именем пакета

```bash
$ go test
ok module3_1-lesson 0.098s
```

Если мы запустим его снова без внесения каких-либо изменений, возвращается кешированный результат без фактического выполнения:

```bash
$ go test
ok module3_1-lesson (cached)
```

Когда тесты такие маленькие, как здесь, на самом деле не имеет значения, кешируются результаты или нет. Но в большом монорепозитории, в котором тысячи тестов в десятках пакетов, попадания и промахи в кеше могут существенно влиять на производительность.

Однако, тесты бывает нужно перезапустить по-настоящему, не изменяя код. Например, если мы обнаружили, что в базе данных нужно внести изменения.

Команда go test не знает ничего про базу данных и выдаст нам кешированный результат. Решить проблему поможет следующая команда

## Команда go clean

Команда go clean удаляет временные файлы, которые были созданы во время сборки проекта. Она также удаляет кеш результатов тестов. Давайте попробуем запустить тесты после очистки кеша

```bash
go clean -cache
go test -v
```

## Флаг -count

Команда go test позволяет запускать тесты несколько раз. Давайте попробуем запустить тесты 3 раза

```bash
go test -v -count=3
```

Зачем это нужно? Есть случаи, когда тест падает не в 100% запусков. Причиной этого может быть то, что тест зависит от внешних факторов, например, от базы данных. Если тест проходит успешно не в 100% запусков и при отсутствии внешних факторов, то это значит что тест написан неверно.

Давайте изменим код программы, чтобы сохранять имена в типе map

```go
package printer

import "fmt"

var names = make(map[string]string)

func PrintHello(name string) string {
    names[name] = name
    return fmt.Sprintf("Hello, %s!", name)
}
```

Запустим тесты 3 раза

```bash
go test -v -count=3
```

Тесты опять проходят успешно:

```bash
   go test -v -count=3

=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.002s
```

Однако, если мы добавим еще одно имя в map names

```go
func TestPrintHello(t *testing.T) {
    names["Petr"] = "Petr"
    got := PrintHello("Igor")
    expected := "Hello, Igor!"

    if got != expected {
        t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
    }

    for name := range names {
        if name != "Igor" {
            t.Fatalf(`got %q, want %q`, name, "Igor")
        } else {
            break
        }
    }
}
```

PrintHello добавляет имя в map names. Встроенный оператор range возвращает элементы в случайном порядке. При каждом запуске теста порядок элементов в map names будет разным. Поэтому тест падает. Несмотря на простоту примера, эта ошибка часто встречается в реальных проектах, даже когда код пишут опытные разработчики, будьте внимательны.

По завершению упражнения откатите изменения в тесте TestPrintHello

## Запуск отдельных тестов

Для запуска отдельных тестов используется опция -run. Она позволяет запускать тесты по имени. Например, если мы хотим запустить только тест TestPrintHello, то мы можем использовать следующую команду:

```bash
go test -v -run=TestPrintHello
```

Мы видим, что тесты проходят успешно:

```bash
go test -v -run=TestPrintHello
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.001s
```

Давайте добавим еще один тест:

```go
func TestPrintHelloIvan(t *testing.T) {
    got := PrintHello("Ivan")
    expected := "Hello, Ivan!"

    if got != expected {
        t.Fatalf(`PrintHello("Ivan") = %q, want %q`, got, expected)
    }
}
```

Запустим только новый тест:

```bash
go test -v -run=TestPrintHelloIvan
```

Видим, что тесты проходят успешно

```bash
go test -v -run=TestPrintHelloIvan
=== RUN   TestPrintHelloIvan
--- PASS: TestPrintHelloIvan (0.00s)
PASS
```

TestPrintHello при этом не запускается.

А теперь давайте выполним следующую команду:

```bash
go test -v -run=TestPrintHello
```

Мы видим, что оба теста запустились:

```bash
go test -v -run=TestPrintHello
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHelloIvan
--- PASS: TestPrintHelloIvan (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.002s
```

Это происходит, потому что TestPrintHelloIvan начинается с TestPrintHello. Так работает -run опция. Поэтому, чтобы не запускать тест TestPrintHelloIvan, когда мы выполняем go test -v -run=TestPrintHelloIvan, мы можем использовать следующую команду:

```bash
go test -v -run=^TestPrintHello$
```

```bash
go test -v -run=^TestPrintHello$
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.002s
```

Это раскрывает механизм работы -run опции. Она принимает регулярное выражение. В нашем случае мы указываем, что тест должен начинаться с TestPrintHello и заканчиваться там же. Таким образом, мы запускаем только тест TestPrintHello.

Но лучше давать уникальные названия тестам, чтобы названия не пересекались. Это потребует меньше усилий и будет проще для работы.

## Опция для запуска всех тестов в пакете

Опция ./... позволяет запускать тесты во всех пакетах внутри текущего модуля. Давайте попробуем запустить тесты во всех пакетах внутри текущего модуля:

```bash
go test -v ./...
```

Мы видим, что тесты проходят успешно:

```bash
go test -v ./...
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHelloIvan
--- PASS: TestPrintHelloIvan (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.001s
```

Разницы между командами go test -v и go test -v ./... нет. Они запускают одни и те же тесты.

Теперь давайте добавим папку counter внутри папки internal и добавим туда 2 файла:

> main.go:

```go
package counter

var counter int

func Increment() {
  counter++
}
```

> main_test.go:

```go
package counter

import "testing"

func TestIncrement(t *testing.T) {
  Increment()

  if counter != 1 {
    t.Error("counter is expected to be incremented")
  }
}
```

Структура директорий должна выглядеть примерно так:

```bash
tree
.
├── go.mod
├── internal
│   └── counter
│       ├── main.go
│       └── main_test.go
├── main.go
└── main_test.go
```

Запустим тесты во всех пакетах внутри текущего модуля:

```bash
go test -v ./...
```

Мы видим, что тесты проходят успешно:

```bash
  go test -v ./...
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHelloIvan
--- PASS: TestPrintHelloIvan (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  (cached)
=== RUN   TestIncrement
--- PASS: TestIncrement (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00/internal/counter 0.001s
```

Запустим go test -v:

```bash
go test -v
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintHelloIvan
--- PASS: TestPrintHelloIvan (0.00s)
PASS
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00  0.001s
```

Мы видим, что тесты в пакете counter не запускаются. Это происходит, потому что go test запускает тесты только в текущем пакете project.

В реальной жизни код проекта разложен по директориям и пакетам. Поэтому важно уметь запускать тесты во всех пакетах внутри текущего модуля.

## Table-driven tests

Table driven testing — термин который говорит, что тест должен состоять из множества случаев. Пример table driven testing:

```go
func TestSum(t *testing.T) {
        // набор тестов
        cases := []struct {
            // имя теста
            name string
            // значения на вход тестируемой функции
            values []int
            // желаемый результат
            want int
        }{
            // тестовые данные №1
            {
                name: "positive values",
                values: []int{1, 2, 3},
                want: 6,
            },
            // тестовые данные №2
            {
                name: "mixed values",
                values: []int{-1, 2, -3},
                want: -2,
            },
        }
        // перебор всех тестов
        for _, tc := range cases {
            tc := tc
            // запуск отдельного теста
            t.Run(tc.name, func(t *testing.T) {
                // тестируем функцию Sum
                got := Sum(tc.values...)
                // проверим полученное значение
                if got != tc.want {
                        t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
                }
            })
        }
}
```

## Test coverage

Test coverage — это термин, который описывает, какая часть кода пакета выполняется при запуске тестов пакета. Если выполнение набора тестов приводит к выполнению 80% кода, мы говорим, что покрытие тестами составляет 80%. Например, если у вас есть пакет с одной функцией:

```go
func Length(a int) string {
    switch {
    case a < 0:
        return "negative"
    case a == 0:
    return "zero"
    case a < 10:
        return "short"
    case a < 100:
        return "long"
    }
    return "very long"
}
```

и тестирующий код вида:

```go
type Test struct {
    in  int
    out string
}

var tests = []Test{
    {-1, "negative"},
    {5, "short"},
}

func TestLength(t *testing.T) {
    for i, test := range tests {
        size := Length(test.in)
        if size != test.out {
            t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
        }
    }
}
```

Теперь, если запустить go test -cover, мы получим следующий результат:

```bash
go test -cover
PASS
coverage: 50.0% of statements
ok      github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-03/01.StandartTestInGo/00/internal/length  0.001s
```

Такой результат мы получаем из-за того, что некоторые участки кода из функции Length не задействованы в тесте.

## Итого

- -v — выводит подробную информацию о тестах
- go clean — очищает кеш результатов тестов
- -count=1 — запускает тесты нужное количество раз
- -run — запускает тесты по имени
- ./... — запускает тесты во всех пакетах внутри текущего модуля
- -cover — выводит информацию о покрытии кода тестами
