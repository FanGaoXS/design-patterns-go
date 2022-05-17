package main

import "fmt"

func main() {
	ak, _ := getGun("ak")
	m4, _ := getGun("m4")

	fmt.Printf("ak = %#v\n", ak)
	fmt.Printf("m4 = %#v\n", m4)
}
