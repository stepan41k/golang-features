package main

type Chair interface {
	HasLegs() bool
}

type Sofa interface {
	Soft() bool
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

type ModernChair struct {}

func (mc ModernChair) HasLegs() bool {
	return false
}

type ModernSofa struct {}

func (ms ModernSofa) Soft() bool {
	return true
}


type ModernFactory struct {}

func (mf ModernFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (mf ModernFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}