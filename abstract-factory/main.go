package main

import "fmt"

func main() {

	adidasFactory, _ := getSportFactory("adidas")
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	nikeFactory, _ := getSportFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

}

func printShoeDetails(s iShoe) {
	fmt.Println("long =", s.Long())
	fmt.Println("price =", s.Price())
}

func printShirtDetails(s iShirt) {
	fmt.Println("height =", s.Height())
	fmt.Println("width =", s.Width())
	fmt.Println("price =", s.Price())
}
