package main

type TreeType struct {
	Name string
	Color string
}

type Tree struct {
	X, Y int
	Type *TreeType
}