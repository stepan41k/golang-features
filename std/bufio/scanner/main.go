package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Создаем сканер для чтения из стандартного потока ввода (клавиатура)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите ваше имя: ")

	// Метод Scan() считывает данные до следующего перевода строки
	if scanner.Scan() {
		name := scanner.Text() // Получаем строку
		fmt.Printf("Привет, %s!\n", name)
	}

	// Проверка на ошибки при чтении
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
	}
}