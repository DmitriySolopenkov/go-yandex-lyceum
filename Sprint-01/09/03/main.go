package main

import (
	"errors"
)

type Account struct {
	owner   string
	balance float64
}

func NewAccount(o string) Account {
	return Account{owner: o}
}

// Метод для установки баланса
func (a *Account) SetBalance(newBalance float64) error {
	if newBalance >= 0 {
		a.balance = newBalance
		return nil
	} else {
		return errors.New("error")
	}
}

// Метод для получения баланса
func (a Account) GetBalance() float64 {
	return a.balance
}

// Метод для внесения денег на счет
func (a *Account) Deposit(amount float64) error {
	if amount > 0 {
		a.balance += amount
		return nil
	} else {
		return errors.New("error")
	}
}

// Метод для снятия денег со счета
func (a *Account) Withdraw(amount float64) error {
	if amount > 0 {
		if amount <= a.balance {
			a.balance -= amount
			return nil
		} else {
			return errors.New("error")
		}
	} else {
		return errors.New("error")
	}
}

/*
Банковский счет
Напишите программу для управления банковским счетом.
Создайте структуру Account с приватными полями balance (баланс)
	и owner (владелец).

Реализуйте методы для установки баланса и получения баланса,
	а также методы для внесения и снятия денег с счета.

Убедитесь, что баланс не может быть отрицательным.

Примечания
Код программы должен содержать описание струкрутры Account:
type Account struct { owner string balance float64 }
*/
