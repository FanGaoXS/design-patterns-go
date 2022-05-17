package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak":
		return newAK(), nil
	case "m4":
		return newM4(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
