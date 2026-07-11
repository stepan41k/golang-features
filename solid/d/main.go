package main

// BAD
// type MySQL struct {}

// func (db MySQL) Query() string {
// 	return "data"
// }

// type App struct {
// 	db MySQL
// }

type Storage interface {
	Query() string
}

type App struct {
	storage Storage
}

func NewApp(s Storage) *App {
	return &App{
		storage: s,
	}
}
