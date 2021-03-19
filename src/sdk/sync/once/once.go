package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for index := 0; index < 10; index++ {
		go func() {
			once.Do(func() {
				fmt.Println("begin to wait")
				time.Sleep(5 * time.Second)
				fmt.Println("once")
			})
		}()
		fmt.Printf("time:%d\n", index)
	}
	time.Sleep(10 * time.Second)
}
