package main

import (
	"fmt"
	"math"
)

func main() {
	var signed int8 = math.MaxInt8
	signed++

	var unsigned uint8 = math.MaxUint8
	unsigned++

	fmt.Println(signed)
	fmt.Println(unsigned)
}