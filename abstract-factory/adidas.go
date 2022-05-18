package main

type adidas struct {
	iSportFactory
}

func (a adidas) makeShirt() iShirt {
	return adidasShirt{
		width:  10,
		height: 10,
		price:  99,
	}
}

func (a adidas) makeShoe() iShoe {
	return adidasShoe{
		long:  10,
		price: 199,
	}
}

type adidasShirt struct {
	iShirt

	width  int
	height int
	price  int
}

func (a adidasShirt) Width() int {
	return a.width
}

func (a adidasShirt) Height() int {
	return a.height
}

func (a adidasShirt) Price() int {
	return a.price
}

type adidasShoe struct {
	iShoe

	long  int
	price int
}

func (a adidasShoe) Long() int {
	return a.long
}

func (a adidasShoe) Price() int {
	return a.price
}
