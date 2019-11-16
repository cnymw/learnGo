package p27

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 2))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
