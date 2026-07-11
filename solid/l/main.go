package main

// BAD

// type Bird interface {
// 	Fly() string
// }

// type Duck struct {}

// func (d Duck) Fly() string {
// 	return "Duck flying"
// }

// type Ostrich struct {}

// func (o Ostrich) Fly() string {
// 	panic("Ostrich cant flying")
// }

// GOOD
type Walker interface {
	Walk()
}

type Flier interface {
	Fly()
}
