package main

type Coffee interface {
	GetPrice() int
}

type SimpleCoffee struct {}

func (c *SimpleCoffee) GetPrice() int {
	return 100
}

type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) GetPrice() int {
	return m.coffee.GetPrice() + 20
}

