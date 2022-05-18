package main

type nike struct {
	iSportFactory
}

func (n nike) makeShirt() iShirt {
	return nikeShirt{
		height: 50,
		width:  50,
		price:  199,
	}
}

func (n nike) makeShoe() iShoe {
	return nikeShoe{
		long:  99,
		price: 399,
	}
}

type nikeShirt struct {
	iShirt

	height int
	width  int
	price  int
}

func (n nikeShirt) Height() int {
	return n.height
}

func (n nikeShirt) Width() int {
	return n.width
}

func (n nikeShirt) Price() int {
	return n.price
}

type nikeShoe struct {
	iShoe

	long  int
	price int
}

func (n nikeShoe) Long() int {
	return n.long
}

func (n nikeShoe) Price() int {
	return n.price
}
