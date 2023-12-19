package main

import "log"

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

type OrderLogger struct{}

func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}

func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)

}

// package main

// type Order struct {
// 	OrderNumber  int
// 	CustomerName string
// 	OrderAmount  float64
// }
// type OrderLogger struct{}

// func NewOrderLogger() *OrderLogger {
// 	return &OrderLogger{}
// }

// func (logger *OrderLogger) AddOrder(order Order) {
// 	logger.l.Info("test")

// }
func main() {
	logger := NewOrderLogger()
	order := Order{1, "Иванов", 100.50}
	logger.AddOrder(order)
}

// package main

// import (
// 	"github.com/sirupsen/logrus"
// )

// type Order struct {
// 	OrderNumber  int
// 	CustomerName string
// 	OrderAmount  float64
// }
// type OrderLogger struct {
// 	Log *logrus.Logger
// }

// func NewOrderLogger() *OrderLogger {
// 	log := logrus.New()
// 	log.SetFormatter(&logrus.JSONFormatter{})
// 	return &OrderLogger{
// 		Log: log,
// 	}
// }

// func (logger *OrderLogger) AddOrder(order Order) {
// 	logger.Log.WithFields(logrus.Fields{
// 		"order_number":  order.OrderNumber,
// 		"customer_name": order.CustomerName,
// 		"order_amount":  order.OrderAmount,
// 	}).Info("New order added")

// }

/*
Логи в терминал
Напишите функцию (logger *OrderLogger) AddOrder(order Order),
	которая пишет в терминал информацию о добавленом заказе вида
	Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f

Примечания
// Order представляет информацию о заказе.
type Order struct {
    OrderNumber  int
    CustomerName string
    OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
    return &OrderLogger{}
}
*/
