package main

import (
	"fmt"
)

type Computer struct {
	CPU string
	RAM int
	GPU string
}

func (c Computer) PrintCPU() {
	fmt.Println(c.CPU)
}

type Builder struct {
	computer Computer
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetCPU(cpu string) *Builder {
	b.computer.CPU = cpu
	return b
}

func (b *Builder) SetRAM(ram int) *Builder {
	b.computer.RAM = ram
	return b
}

func (b *Builder) Build() Computer {
	return b.computer
}

func main() {
	pc := NewBuilder().SetCPU("i9").SetRAM(16).Build()
	pc.PrintCPU()
}