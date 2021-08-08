package p15

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}

func TestThreeSum1(t *testing.T) {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
