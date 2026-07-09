package main

type Subject interface {
	Subscribe(Observer)
	NotifyAll()
}

type Item struct {
	observers []Observer
	inStock bool
}

func (i *Item) UpdateAvailabillity() {
	i.inStock = true
	i.NotifyAll()
}