package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine3(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}
func Routine1(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		value++
		Counter = value
	}

	Wait.Done()
}
func Routine2(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
	}

	Wait.Done()
}
func Routine3(id int) {

	for count := 0; count < 2; count++ {

		Counter = Counter + 1
		time.Sleep(1 * time.Nanosecond)
	}

	Wait.Done()
}
