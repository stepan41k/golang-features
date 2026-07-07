package main

import (
	"fmt"
	"sync"
)

type database struct {}

var (
	instance *database
	once sync.Once
)

func GetDatabaseInstance() *database {
	once.Do(func() {
		fmt.Println("Создание экземпляра БД...")
		instance = &database{}
	})

	return instance
}