package main

import "fmt"

func computePrefix(pattern string) []int {
	// Создай массив pi длиной как паттерн.
	// Используй два указателя: j (длина текущего префикса) и i (текущий символ, начиная с 1).
	// Если P[i] == P[j], то j++, записываем pi[i] = j.
	// Если не равны и j > 0, откатываем j к pi[j-1] и пробуем снова.
	pi := make([]int, len(pattern))
	j, i := 0, 1

	if pattern[i] == pattern[j] {
		j++
		pi[i] = j
	} else {
		if j > 0 {
			j = pi[j-1]
		}
	}

	return pi
}

func KMP(text, pattern string) []int {
	res := []int{}
	fmt.Println(computePrefix(pattern))
	// Мы идем по тексту T один раз. Если символы совпадают — идем дальше. Если нет — смотрим в таблицу pi, чтобы узнать, на какую позицию в паттерне мы можем откатиться, не перепроверяя уже совпавшие символы в тексте.

	// Возвращает слайс всех индексов (начал) вхождений паттерна в текст.

	return res
}

func main() {
	KMP("ABCDABD", "ABCAD")
}