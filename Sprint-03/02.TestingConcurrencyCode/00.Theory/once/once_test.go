package once

import (
	"sync"
	"testing"
)

func TestIncrementOnce(t *testing.T) {
	// Сбросим счётчик
	counter = 0

	// Функция для увеличения значения счётчика
	IncrFunc = func() {
		counter++
	}

	// Для ожидания горутин
	var wg sync.WaitGroup

	// 	// Будем делать инкремент столько раз
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
		t.Errorf(
			"Expected counter value: %d, Actual counter value: %d",
			expectedCounter,
			actualCounter,
		)
	}
}
