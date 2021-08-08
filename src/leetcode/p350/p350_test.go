package p350

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2}))

}
