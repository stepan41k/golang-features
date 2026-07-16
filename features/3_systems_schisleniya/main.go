package main

import "fmt"

func main() {
	// bin - 0b OR 0B (f.e. 0b100)
	x2 := 0b100

	// oct - 0 OR 0o (f.e. 010)
	x8 := 010

	// hex - 0x OR 0X (f.e. 0xF)
	x16 := 0xF

	// mnim - i (f.e. 3i)
	xm := 3i

	x10 := 100_000_000

	fmt.Println(x2, x8, x16, xm, x10)
}