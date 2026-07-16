package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var y int8 = 3
	var x uint8 = 3

	fmt.Println(^y)
	fmt.Println(^x)

	var y2 int8 = 1

	fmt.Println(y2 << 6)
	fmt.Println(y2 << 7)
}