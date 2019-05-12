package tool

import (
	"fmt"
	"learnGo/src/algorithm/sort/sort_interface"
	"math/rand"
	"reflect"
	"time"
)

// 生成大小为10的随机数组
func GenerateRandomList() ([10]int, error) {
	rand.Seed(time.Now().UnixNano())
	list := [10]int{}
	for i := 0; i < 10; i++ {
		list[i] = rand.Intn(100)
	}
	return list, nil
}

// 使用策略模式来调用各个排序算法
func SortTool(s sortinterface.Interface) {
	// 利用反射获取接口的名称
	defer trace(reflect.ValueOf(s).String())()
	fmt.Printf("before sort : %v\n", s)
	s.Sort()
	fmt.Printf("after sort : %v\n", s)
}

// 通过defer函数来记录算法的运行时间
func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)\n\n", msg, time.Since(start))
	}
}
