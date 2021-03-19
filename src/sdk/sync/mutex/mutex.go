package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

var m sync.Mutex

func main() {
	lockSlow1()
	//lockSlow2()
	//
	//time.Sleep(5 * time.Second)
}

func race() {
	var x int
	go func() {
		for {
			x = 1
			fmt.Println(x)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			x = 2
			fmt.Println(x)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			x = 3
			fmt.Println(x)
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}

func lockSlow1() {
	//old := 0
	fmt.Println(1 & 1)
}

func lockSlow2() {
	go func() {
		for {
			m.Lock()
			fmt.Println("get lock 1")
			time.Sleep(time.Second)
			m.Unlock()
		}
	}()

	go func() {
		for {
			m.Lock()
			fmt.Println("get lock 2")
			time.Sleep(time.Second)
			m.Unlock()
		}
	}()
}
