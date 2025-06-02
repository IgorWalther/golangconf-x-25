package main

import (
	"fmt"
)

var value string
var done = make(chan struct{})

func setup() {
	value = "Hello, GolangConf-X-2025!" // (3)
	done <- struct{}{}                  // (4)
}

func main() {
	go setup()
	<-done //(1)

	fmt.Println(value) // (2)
}
