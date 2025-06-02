package main

import (
	"fmt"
	"sync/atomic"
)

var value string
var done atomic.Bool

func setup() {
	value = "Hello, GolangConf-X-2025!" // (3)
	done.Store(true)                    // (4)
}

func main() {
	go setup()

	for !done.Load() {
	} // (1)

	fmt.Println(value) // (2)
}
