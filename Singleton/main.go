package main

import "time"

func main() {
	getDatabaseInstance("test")
	for i := 0; i < 10; i++ {
		go getDatabaseInstance("test1")
		time.Sleep(time.Second * 1)
	}
}
