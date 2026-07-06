package main

import (
	"fmt"
)

type Transport interface {
	Deliver() string
}

type Truck struct {}

func (t Truck) Deliver() string {
	return "Доставка по земле"
}

type Ship struct {}

func (t Ship) Deliver() string {
	return "Доставка по воде"
}

func NewTransport(mType string) (Transport, error) {
	if mType == "truck" {
		return &Truck{}, nil
	}

	if mType == "ship" {
		return &Ship{}, nil
	}

	return nil, fmt.Errorf("unknown transport type")
}

func main() {
	t, _ := NewTransport("truck")
	fmt.Println(t.Deliver())
}
