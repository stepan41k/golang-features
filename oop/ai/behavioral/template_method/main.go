package main

import "fmt"

type Notifier interface {
	GetMessage() string
	Send(msg string)
}

// "Шаблонный" процесс
func Notify(n Notifier) {
	msg := n.GetMessage()
	n.Send(msg)
	fmt.Println("Уведомление отправлено")
}

type SMSNotifier struct {}

func (s SMSNotifier) GetMessage() string {
	return "Привет по SMS"
}

func (s SMSNotifier) Send(m string) {
	fmt.Println("Отправка по SMS:", m)
}