package main

import (
	"fmt"
	"time"
)

var v int

func g0() {
	for {
		v = 0
		fmt.Println(v)
	}
}

func g1() {
	for {
		v = 1
		fmt.Println(v)
	}
}

func main() {
	go g0()
	go g1()

	time.Sleep(time.Minute)
}
