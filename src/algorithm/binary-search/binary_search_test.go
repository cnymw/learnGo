package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 4, 10, 40}

	fmt.Println(binarySearchRecursive(arr, 0, len(arr)-1, 3))

	fmt.Println(binarySearchIterative(arr, 0, len(arr)-1, 3))
}
