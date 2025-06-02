package main

import (
	"fmt"
	"sync"
)

var (
	value string
	done  bool
	mx    sync.RWMutex
)

func setup() {
	value = "Hello, GolangConf-X-2025!" // (3)
	mx.Lock()
	done = true // (4)
	mx.Unlock()
}

func main() {
	go setup()

	for {
		mx.RLock()
		snapshot := done // (1)
		mx.RUnlock()

		if snapshot {
			break
		}
	}
	fmt.Println(value) // (2)
}
