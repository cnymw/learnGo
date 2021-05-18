package main

import (
	"fmt"
	"log"
	"time"
)

type Obj struct {
	val int
}

func main() {
	for index := 0; index < 100; index++ {
		go func(index int) {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
			var obj *Obj
			fmt.Println(index)
			if index%3 == 0 {
				fmt.Println(obj.val)
			}
		}(index)
	}

	time.Sleep(10 * time.Second)
}
