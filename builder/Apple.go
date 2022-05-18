package main

type AppleBuilder struct {
	cpu     string
	screen  string
	battery string
}

func NewAppleBuilder() *AppleBuilder {
	return &AppleBuilder{}
}

func (a *AppleBuilder) BuildCPU(cpu string) {
	a.cpu = cpu
}

func (a *AppleBuilder) BuildScreen(screen string) {
	a.screen = screen
}

func (a *AppleBuilder) BuildBattery(battery string) {
	a.battery = battery
}

func (a AppleBuilder) GetPhone() Phone {
	return Phone{
		cpu:     a.cpu,
		screen:  a.screen,
		battery: a.battery,
	}
}
