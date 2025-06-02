package main

import (
	"fmt"
	"time"
)

var value string

func doPrint() {
	fmt.Println(value) // (3)
}

func main() {
	value = "Hello, GolangConf!" // (1)
	go doPrint()                 // (2)

	time.Sleep(time.Minute)
}
