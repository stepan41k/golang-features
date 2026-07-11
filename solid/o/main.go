package main

import (
	"fmt"
)

// BAD
// func ProccessPayment(amount int, method string) {
// 	if method == "paypal" {}
// 	if method == "card" {}
// }

// GOOD
type PaymentMethod interface {
	Pay(amount int)
}

type CreditCard struct {}
func (c CreditCard) Pay(amount int) {
	fmt.Println("Оплата картой")
}

type Crypto struct {}
func (c Crypto) Pay(amount int) {
	fmt.Println("Оплата биткоином")
}

func Proccess(amount int, m PaymentMethod) {
	m.Pay(amount)
}

