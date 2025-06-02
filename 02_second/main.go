package main

import (
	"fmt"
)

var value string
var done bool

func setup() {
	value = "Hello, GolangConf-X-2025!" // (3)
	done = true                         // (4)
}

func main() {
	go setup()

	for !done {
	} // (1)

	fmt.Println(value) // (2)
}
