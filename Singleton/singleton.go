package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var singletonInstance *singleton

func getSingletonInstance() *singleton {
	if singletonInstance == nil {
		once := sync.Once{}
		once.Do(func() {
			fmt.Println("create instance now")
			singletonInstance = &singleton{}
		})
	} else {
		fmt.Println("instance already created")
	}
	return singletonInstance
}
