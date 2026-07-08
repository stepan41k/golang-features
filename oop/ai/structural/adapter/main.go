package main

import "fmt"

type ModernComputer interface {
	InsertInSquarePort()
}

type OldUSB struct {}

func (u *OldUSB) InsertInRoundPort() {
	fmt.Println("USB-A вставлен")
}

type USBAdapter struct {
	oldUSB *OldUSB
}

func (a *USBAdapter) InsertInSquarePort() {
	fmt.Println("Адаптер преобразует сигнал")
	a.oldUSB.InsertInRoundPort()
}