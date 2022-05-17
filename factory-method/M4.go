package main

type M4 struct {
	name  string
	power int
}

func newM4() iGun {
	return &M4{
		name:  "m4a1",
		power: 1,
	}
}

func (m M4) getName() string {
	return m.name
}

func (m M4) getPower() int {
	return m.power
}
