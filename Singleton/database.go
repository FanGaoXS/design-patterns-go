package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type database struct {
	addr string
}

var databaseInstance *database

func getDatabaseInstance(addr string) *database {
	// 如果实例不存在，说明实例还没有被初始化过，则执行实例的初始化
	lock.Lock()
	defer lock.Unlock()
	if databaseInstance == nil {
		fmt.Println("create instance now")
		databaseInstance = &database{addr: addr}
	} else {
		fmt.Println("instance already created")
	}
	return databaseInstance
}
