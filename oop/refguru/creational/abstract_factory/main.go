package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShirtDetails(nikeShoe)
	printShoeDetails(nikeShirt)

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(s IShoe) {
    fmt.Printf("Logo: %s\n", s.getLogo())
    fmt.Printf("Size: %d\n", s.getSize())
    fmt.Printf("\n")
}

func printShirtDetails(s IShirt) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}