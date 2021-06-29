package p1

import (
	"fmt"
	"testing"
)

func TestTwoSumByLoop(t *testing.T) {
	fmt.Println(twoSumByLoop([]int{3, 2, 4}, 6))
}

func TestTwoSumByHash(t *testing.T) {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumByHashMap([]int{3, 2, 4}, 6))
}
