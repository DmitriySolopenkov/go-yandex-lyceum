package main

import (
	"testing"
	"time"
)

func TestCreateReport(t *testing.T) {
	user := User{ID: 1, Name: "Ğ˜Ğ²Ğ°Ğ½", Email: "ivan@example.com", Age: 30}
	reportDate := time.Now().Format("2006-01-02")

	report := CreateReport(user, reportDate)

	// ĞŸÑ€Ğ¾Ğ²ĞµÑ€Ñ�ĞµĞ¼, Ñ‡Ñ‚Ğ¾ Ğ¿Ğ¾Ğ»Ñ� Ğ² Ğ¾Ñ‚Ñ‡ĞµÑ‚Ğµ Ğ·Ğ°Ğ¿Ğ¾Ğ»Ğ½ĞµĞ½Ñ‹ ĞºĞ¾Ñ€Ñ€ĞµĞºÑ‚Ğ½Ğ¾
	if report.ID != user.ID {
		t.Errorf("Ğ�Ğ¶Ğ¸Ğ´Ğ°ĞµÑ‚Ñ�Ñ� ID Ğ¿Ğ¾Ğ»ÑŒĞ·Ğ¾Ğ²Ğ°Ñ‚ĞµĞ»Ñ� %d, Ğ¿Ğ¾Ğ»ÑƒÑ‡ĞµĞ½Ğ¾ %d", user.ID, report.ID)
	}
	if report.Name != user.Name {
		t.Errorf("Ğ�Ğ¶Ğ¸Ğ´Ğ°ĞµÑ‚Ñ�Ñ� Ğ¸Ğ¼Ñ� Ğ¿Ğ¾Ğ»ÑŒĞ·Ğ¾Ğ²Ğ°Ñ‚ĞµĞ»Ñ� %s, Ğ¿Ğ¾Ğ»ÑƒÑ‡ĞµĞ½Ğ¾ %s", user.Name, report.Name)
	}
	if report.Email != user.Email {
		t.Errorf("Ğ�Ğ¶Ğ¸Ğ´Ğ°ĞµÑ‚Ñ�Ñ� Email Ğ¿Ğ¾Ğ»ÑŒĞ·Ğ¾Ğ²Ğ°Ñ‚ĞµĞ»Ñ� %s, Ğ¿Ğ¾Ğ»ÑƒÑ‡ĞµĞ½Ğ¾ %s", user.Email, report.Email)
	}
	if report.Age != user.Age {
		t.Errorf("Ğ�Ğ¶Ğ¸Ğ´Ğ°ĞµÑ‚Ñ�Ñ� Ğ²Ğ¾Ğ·Ñ€Ğ°Ñ�Ñ‚ Ğ¿Ğ¾Ğ»ÑŒĞ·Ğ¾Ğ²Ğ°Ñ‚ĞµĞ»Ñ� %d, Ğ¿Ğ¾Ğ»ÑƒÑ‡ĞµĞ½Ğ¾ %d", user.Age, report.Age)
	}
}
