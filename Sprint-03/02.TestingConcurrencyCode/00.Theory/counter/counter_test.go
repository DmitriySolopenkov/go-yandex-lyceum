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
		t.Errorf(
			"Expected counter value: %d, Actual counter value: %d",
			expectedCounter,
			actualCounter,
		)
	}
}
