package main

type AK struct {
	name  string
	power int
}

func newAK() iGun {
	return &AK{
		name:  "ak47",
		power: 5,
	}
}

func (a AK) getName() string {
	return a.name
}

func (a AK) getPower() int {
	return a.power
}
