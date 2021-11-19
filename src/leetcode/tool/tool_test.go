package tool

import (
	"fmt"
	"testing"
)

func TestArrayToTree(t *testing.T) {
	array := []int{3, 9, 20, 0, 0, 15, 7}

	fmt.Println(ArrayToTree(array).Val)
}
