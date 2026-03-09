package main

import (
	"fmt"
)

// Реализуй функцию setBit(n int, i int) int.
// Функция должна принимать число n и устанавливать его i-ый бит в 1, оставляя остальные биты без изменений.
// Пример:
// Вход: n = 8 (1000₂), i = 0
// Результат: 9 (1001₂)

func setBit(x int, i int) int {
	return x | (1 << i)			// 1000 + 0001 = 1001
}

func main() {
	fmt.Println(setBit(8, 4))
}
