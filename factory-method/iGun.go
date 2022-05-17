package main

type iGun interface {
	getName() string // 返回武器的名称
	getPower() int   // 返回武器的能量大小
}
