package p11

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println(maxArea([]int{1, 1}))
}
