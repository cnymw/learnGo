package tool

import (
	"fmt"
	"learnGo/src/algorithm/sort/sort_interface"
	"math/rand"
	"time"
)

func GenerateRandomList() ([10]int, error) {
	rand.Seed(time.Now().UnixNano())
	list := [10]int{}
	for i := 0; i < 10; i++ {
		list[i] = rand.Intn(100)
	}
	return list, nil
}

func SortTool(s sortinterface.Interface) {
	fmt.Printf("before sort : %v\n", s)
	s.Sort()
	fmt.Printf("after sort : %v\n", s)
}
