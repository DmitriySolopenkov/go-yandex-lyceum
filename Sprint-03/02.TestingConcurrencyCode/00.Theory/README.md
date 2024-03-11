# Тестирование concurrency кода

## Введение

Тестирование конкурентного кода важно для обеспечения корректной работы программы в условиях параллелизма. В этом уроке мы рассмотрим основы тестирования конкурентного кода в языке программирования Go.

## Практические аспекты тестирования конкурентного кода

Давайте рассмотрим несколько практических аспектов тестирования конкурентного кода:

### Использование t.Parallel()

В Go, тесты могут быть запущены параллельно для проверки работы программы в паралельном режиме и для ускорения их выполнения, если тесты не связанные. Если ваша программа или часть программы, которую вы тестируете не зависит от незащищенных глобальных переменных или ресурсов, добавьте t.Parallel() в начало теста, чтобы разрешить его параллельное выполнение.

Вы можете встретиться с требованием использования t.Parallel() во всех тестах в некоторых компаниях.

Пример:

```go
func TestIncrementParallel(t *testing.T) {
t.Parallel()

// тело теста...
}
```

### Использование sync.WaitGroup для ожидания завершения горутин

В тестах, где запускаются горутины, используйте sync.WaitGroup, чтобы дождаться их завершения перед завершением теста.

Пример:

```go
func TestConcurrentProcessing(t *testing.T) {
var wg sync.WaitGroup

// ... инициализация ...

// Запуск горутин
for i := 0; i < numGoroutines; i++ {
wg.Add(1)
go func (i int) {
defer wg.Done()
// логика горутины...
}(i)
}

// Ожидание завершения всех горутин
wg.Wait()

// ... проверки результатов ...
}
```

Ниже мы используем этот подход в примерах.

Эти практические аспекты помогут вам создавать надёжные и эффективные тесты для конкурентного кода в Go.

## Пример кода 1

Тестирование конкурентного кода включает в себя проверку корректности выполнения кода в условиях одновременного выполнения нескольких горутин. Рассмотрим примеры.

### Тестирование конкурентного кода в Go с использованием t.Parallel

В этом примере мы рассмотрим тестирование конкурентного кода, который выполняет операции добавления и удаления записи из кеша. Мы используем функцию t.Parallel для запуска тестов параллельно.

```go
// cache.go

package cache

import (
    "sync"
)

// Cache представляет собой простой кеш
type Cache struct {
    mu     sync.Mutex
    values map[string]string
}

// NewCache создает новый экземпляр кеша
func NewCache() *Cache {
    return &Cache{
        values: make(map[string]string),
    }
}

// Add добавляет запись в кеш
func (c *Cache) Add(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.values[key] = value
}

// Delete удаляет запись из кеша
func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.values, key)
}

// Get возвращает значение записи по ключу из кеша
func (c *Cache) Get(key string) string {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.values[key]
}
```

Теперь давайте напишем тест для этого кода с использованием t.Parallel

```go
// cache_test.go

package cache

import (
    "testing"
)

// TestConcurrentCache - тест для параллельного тестирования кеша
func TestConcurrentCache(t *testing.T) {
    // Создаем новый кеш
    cache := NewCache()

    // Добавляем запись в кеш (не в параллели)
    cache.Add("key1", "value1")

    // Запускаем тесты в параллели
    t.Run("AddRecord", func(t *testing.T) {
        t.Parallel()
        // Добавляем запись в кеш
        cache.Add("key2", "value2")
    })

    t.Run("DeleteRecord", func(t *testing.T) {
        t.Parallel()
        // Удаляем запись из кеша
        cache.Delete("key1")
    })

    // Проверяем наличие записи после всех операций
    t.Run("CheckRecord", func(t *testing.T) {
        t.Parallel()
        // Получаем значение записи по ключу
        value := cache.Get("key2")
        expectedValue := "value2"

        // Проверяем, что значение соответствует ожидаемому
        if value != expectedValue {
            t.Errorf("Expected value: %s, Actual value: %s", expectedValue, value)
        }
    })
}
```

В этом тесте мы создаём экземпляр кеша и выполняем несколько операций (добавление, удаление) параллельно. Используя t.Parallel, мы можем запускать эти тесты одновременно, что помогает обнаруживать проблемы в конкурентном коде.

## Пример кода 2

### Тестирование инкрементации

Предположим, у нас есть функция, которая инкрементирует общий счётчик:

```go
// counter.go

package counter

import "sync"

var mu sync.Mutex
var counter int

func Increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}

func GetCounter() int {
    mu.Lock()
    defer mu.Unlock()
    return counter
}
```

Теперь давайте напишем тесты для этого кода.

```go
// counter_test.go

package counter

import (
    "sync"
    "testing"
)

func TestIncrement(t *testing.T) {
    // Сбросим счётчик
    counter = 0

    // Для ожидания горутин
    var wg sync.WaitGroup

    // Будем делать инкремент столько раз
    numIncrements := 1000

    // Для ожидания всех запущенных горутин
    wg.Add(numIncrements)

    // Увеличиваем значение счётчика конкурентно
    for i := 0; i < numIncrements; i++ {
        go func() {
            defer wg.Done()
            Increment()
        }()
    }

    // Подождём все горутины
    wg.Wait()

    // Проверим, получили ли ожидаемое значение
    expectedCounter := numIncrements
    actualCounter := GetCounter()

    if actualCounter != expectedCounter {
        t.Errorf("Expected counter value: %d, Actual counter value: %d", expectedCounter, actualCounter)
    }
}
```

Этот тест создает 1000 горутин, каждая из которых вызывает функцию Increment. После завершения всех горутин, тест проверяет, что значение счётчика соответствует ожидаемому.

## Пример кода 3

### Тестирование sync.Once

Давайте рассмотрим еще один пример тестирования конкурентного кода, на этот раз с использованием sync.Once. Представим, что у нас есть функция initialize с использованием sync.Once, которая должна быть вызвана только один раз:

```go
// once.go

package once

import (
    "sync"
)

var (
    once     sync.Once
    counter  int
    incrFunc func()
)

// initialize увеличивает счётчик только один раз
func initialize() {
    counter++
}

// Increment вызывает initialize только один раз
func Increment() {
    once.Do(initialize)
}

// GetCounter возвращает текущее значение счётчика
func GetCounter() int {
    return counter
}
```

Теперь давайте напишем тесты для проверки корректности использования sync.Once

```go
// once_test.go

package once

import (
    "sync"
    "testing"
)

func TestIncrementOnce(t *testing.T) {
    // Сбросим счётчик
    counter = 0

    // Функция для увеличения значения счётчика
    incrFunc = func() {
        counter++
    }

    // Для ожидания горутин
    var wg sync.WaitGroup

    //     // Будем делать инкремент столько раз
    numIncrements := 1000

    // Для ожидания всех запущенных горутин
    wg.Add(numIncrements)

    // Увеличиваем значение счётчика конкурентно и используя sync.Once
    for i := 0; i < numIncrements; i++ {
        go func() {
            defer wg.Done()
            Increment()
        }()
    }

    // Подождём все горутины
    wg.Wait()

    // Проверим, получили ли ожидаемое значение
    expectedCounter := 1 // 1 - потому что использовали sync.Once
    actualCounter := GetCounter()

    if actualCounter != expectedCounter {
        t.Errorf("Expected counter value: %d, Actual counter value: %d", expectedCounter, actualCounter)
    }
}
```

В этом тесте мы используем sync.Once для гарантии того, что функция initialize вызывается только один раз, даже при параллельных вызовах. Такой подход позволяет обеспечить безопасную инициализацию, а тест проверяет, что значение счётчика соответствует ожидаемому результату после всех инкрементов.

## Заключение

Тестирование конкурентного кода важно для обеспечения корректной работы программ в условиях параллелизма. Мы рассмотрели базовые принципы тестирования конкурентного кода в Go, включая использование sync.Mutex для защиты общих ресурсов и написание тестов с использованием горутин и sync.WaitGroup. Эти подходы позволяют обнаруживать потенциальные проблемы в конкурентном коде и обеспечивают стабильность при одновременном выполнении различных частей программы.
