package main

import "fmt"

type Command interface {
	Execute()
}

type SimpleCommand struct {
	payload string
}

func (c *SimpleCommand) Execute() {
	fmt.Println(c.payload)
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}