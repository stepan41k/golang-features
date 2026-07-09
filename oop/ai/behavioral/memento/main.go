package main

type Memento struct {
	state string
}

type Originator struct {
	state string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{
		state: e.state,
	}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.state = m.state
}