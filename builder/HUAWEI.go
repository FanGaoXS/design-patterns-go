package main

type HUAWEIBuilder struct {
	cpu     string
	screen  string
	battery string
}

func NewHUAWEIBuilder() *HUAWEIBuilder {
	return &HUAWEIBuilder{}
}

func (h *HUAWEIBuilder) BuildCPU(cpu string) {
	h.cpu = cpu
}

func (h *HUAWEIBuilder) BuildScreen(screen string) {
	h.screen = screen
}

func (h *HUAWEIBuilder) BuildBattery(battery string) {
	h.battery = battery
}

func (h HUAWEIBuilder) GetPhone() Phone {
	return Phone{
		cpu:     h.cpu,
		screen:  h.screen,
		battery: h.battery,
	}
}
