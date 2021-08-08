package p976

import (
	"fmt"
	"testing"
)

func TestLargestPerimeter(t *testing.T) {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
	fmt.Println(largestPerimeter([]int{2, 6, 2, 5, 4, 15, 1}))
}
