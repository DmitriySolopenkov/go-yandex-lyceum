package main

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser("Тестовый пользователь", 35)

	if user.Name != "Тестовый пользователь" {
		t.Errorf("Ожидается имя 'Тестовый пользователь', получено '%s'", user.Name)
	}

	if user.Age != 35 {
		t.Errorf("Ожидается возраст 35, получено %d", user.Age)
	}

	if user.Active != true {
		t.Errorf("Ожидается активный статус true, получено %v", user.Active)
	}
}
