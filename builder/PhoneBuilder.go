package main

import "fmt"

type PhoneBuilder interface {
	BuildCPU(cpu string)
	BuildScreen(screen string)
	BuildBattery(battery string)
	GetPhone() Phone
}

func GetPhoneBuilder(brand string) (PhoneBuilder, error) {
	switch brand {
	case "Apple":
		return NewAppleBuilder(), nil
	case "HUAWEI":
		return NewHUAWEIBuilder(), nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

type Phone struct {
	cpu     string
	screen  string
	battery string
}
