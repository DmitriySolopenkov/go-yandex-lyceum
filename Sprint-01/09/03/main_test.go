package main

import (
	"testing"
)

func TestAccountOperations(t *testing.T) {
	account := NewAccount("Alice")

	// Устанавливаем баланс
	err := account.SetBalance(1000.0)
	if err != nil {
		t.Errorf("Ошибка при установке баланса: %v", err)
	}

	// Вносим деньги
	err = account.Deposit(500.0)
	if err != nil {
		t.Errorf("Ошибка при внесении денег: %v", err)
	}

	// Снимаем деньги
	err = account.Withdraw(200.0)
	if err != nil {
		t.Errorf("Ошибка при снятии денег: %v", err)
	}

	// Получаем текущий баланс
	balance := account.GetBalance()
	expectedBalance := 1300.0
	if balance != expectedBalance {
		t.Errorf("Ожидается баланс %.2f, получено %.2f", expectedBalance, balance)
	}

	// Пытаемся установить отрицательный баланс
	err = account.SetBalance(-500.0)
	if err == nil {
		t.Errorf("Ожидается ошибка при установке отрицательного баланса, но она не произошла")
	}
}
