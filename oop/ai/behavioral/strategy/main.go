package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int)
}

type PayPal struct {}

func (p *PayPal) Pay(amount int) {
	fmt.Printf("Оплачено %d через PayPal\n", amount)
}

type Checkount struct {
	strategy PaymentStrategy
}

func (c *Checkount) SetStrategy(s PaymentStrategy) {
	c.strategy = s
}

func (c *Checkount) ExecutePayment(amt int) {
	c.strategy.Pay(amt)
}