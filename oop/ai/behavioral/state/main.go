package main

type State interface {
	AddItem(int) error
	RequestItem() error
	DispenseItem() error
}

type VendingMachine struct {
	hasItem State
	noItem State
	currentState State
}

func (v *VendingMachine) RequestItem() {
	v.currentState.RequestItem()
}