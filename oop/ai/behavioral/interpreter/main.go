package main

import "fmt"

type Expression interface {
	Interpret() int
}

type Number struct {
	val int
}

func (n *Number) Interpret() int {
	return n.val
}

type Add struct {
	left, right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

func main() {
	expr := &Add{left: &Number{1}, right: &Number{5}}
	fmt.Println(expr.Interpret())
}